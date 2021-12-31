package main

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	t "github.com/brimstone/go-twitter"
	"github.com/brimstone/logger"
	twitter "github.com/dghubble/go-twitter/twitter"
	"gopkg.in/yaml.v2"
)

type Member struct {
	Description  string
	ID           int64
	Name         string
	ProfileImage string
	ScreenName   string
	Status       *twitter.Tweet
	LastTweet    string
}

type Config struct {
	Lists []string `yaml:"lists"`
}
type TemplateList struct {
	ID      int64
	Name    string
	Members []Member
}

func main() {
	log := logger.New()
	now := time.Now()
	var templists []TemplateList

	var config Config
	configBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	client, _, err := t.NewClient(t.Tokens{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		AccessToken:    os.Getenv("ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("ACCESS_SECRET"),
	})
	if err != nil {
		panic(err)
	}
	listsResp, _, err := client.Lists.List(nil)
	if err != nil {
		panic(err)
	}

	lists := make(map[string]int64)
	for _, list := range listsResp {
		lists[list.Name] = list.ID
		log.Debug("list",
			log.Field("name", list.Name),
			log.Field("ID", list.ID),
		)
	}

	for _, list := range config.Lists {
		membersResp, _, err := client.Lists.Members(&twitter.ListsMembersParams{
			ListID: lists[list],
			Count:  1000,
		})
		if err != nil {
			panic(err)
		}

		log.Debug("members",
			log.Field("count", len(membersResp.Users)),
		)

		members := []Member{}
		for _, member := range membersResp.Users {
			if member.Entities != nil {
				for _, u := range member.Entities.Description.Urls {
					member.Description = strings.ReplaceAll(member.Description, u.URL, u.ExpandedURL)
				}
			}
			m := Member{
				Description:  member.Description,
				ID:           member.ID,
				Name:         member.Name,
				ProfileImage: strings.Replace(member.ProfileImageURLHttps, "_normal", "_200x200", 1),
				ScreenName:   member.ScreenName,
				Status:       member.Status,
			}
			if member.Status != nil {
				c, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", member.Status.CreatedAt)
				if err == nil {
					thisyear, thisweek := now.ISOWeek()
					thatyear, thatweek := c.ISOWeek()
					// TODO Does this handle year wrap around right?
					if thisyear != thatyear {
						thisweek += 52
					}

					// Years
					if now.Year() != c.Year() {
						m.LastTweet = strconv.Itoa(c.Year())
						// Months
						// TODO might error on first day of month
					} else if now.Month() != c.Month() {
						m.LastTweet = c.Month().String()
						// Weeks
					} else if thisweek-thatweek > 1 {
						m.LastTweet = strconv.Itoa(thisweek-thatweek) + " weeks ago"
					} else if thisweek-thatweek == 1 {
						m.LastTweet = "last week"
						/*
								// Days
							} else if now.YearDay()-c.YearDay() > 1 {
								m.LastTweet = strconv.Itoa(now.YearDay()-c.YearDay()) + " days ago"
							} else if now.YearDay()-c.YearDay() == 1 {
								m.LastTweet = "yesterday"
							} else {
								m.LastTweet = "today"
							}
						*/
					} else {
						m.LastTweet = "this week"
					}
				}
			}
			log.Debug("member",
				log.Field("name", member.Name),
				log.Field("ID", member.ID),
				log.Field("screenname", member.ScreenName),
				//log.Field("profileImage", member.ProfileImageURLHttps),
			)
			members = append(members, m)
		}

		sort.Slice(members, func(i, j int) bool {
			return members[i].ID > members[j].ID
		})
		templists = append(templists, TemplateList{
			ID:      lists[list],
			Name:    list,
			Members: members,
		})
	}

	tpl, err := ioutil.ReadFile("README.md.tpl")
	if err != nil {
		panic(err)
	}
	t, err := template.New("readme").Parse(string(tpl))
	if err != nil {
		panic(err)
	}

	data := struct {
		Lists       []TemplateList
		LastUpdated time.Time
	}{
		Lists:       templists,
		LastUpdated: now,
	}

	f, err := os.Create("README.md")
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

}
