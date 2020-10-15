package goGoogleAuth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	/* "google.golang.org/api/classroom/v1" */)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := TokenFromFile(tokFile)
	if err != nil {
		tok = GetTokenFromWeb(config)
		SaveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func SaveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Do() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/classroom.courses https://www.googleapis.com/auth/spreadsheets https://www.googleapis.com/auth/classroom.coursework.students https://www.googleapis.com/auth/classroom.rosters https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/classroom.coursework.me https://www.googleapis.com/auth/classroom.profile.emails https://www.googleapis.com/auth/calendar https://www.googleapis.com/auth/calendar.events https://www.googleapis.com/auth/drive https://www.googleapis.com/auth/drive.appdata https://www.googleapis.com/auth/drive.file https://www.googleapis.com/auth/drive.metadata https://www.googleapis.com/auth/drive.scripts https://www.googleapis.com/auth/drive.activity https://mail.google.com/ https://www.googleapis.com/auth/gmail.addons.current.action.compose https://www.googleapis.com/auth/gmail.addons.current.message.action https://www.googleapis.com/auth/gmail.addons.current.message.metadata https://www.googleapis.com/auth/gmail.compose https://www.googleapis.com/auth/gmail.insert https://www.googleapis.com/auth/gmail.labels https://www.googleapis.com/auth/gmail.metadata https://www.googleapis.com/auth/gmail.modify https://www.googleapis.com/auth/gmail.send https://www.googleapis.com/auth/gmail.settings.basic https://www.googleapis.com/auth/classroom.announcements https://www.googleapis.com/auth/classroom.courseworkmaterials https://www.googleapis.com/auth/classroom.guardianlinks.students https://www.googleapis.com/auth/classroom.push-notifications https://www.googleapis.com/auth/classroom.topics https://www.googleapis.com/auth/documents https://www.googleapis.com/auth/presentations https://www.googleapis.com/auth/contacts https://www.googleapis.com/auth/user.emails.read https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/tasks")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config)
	fmt.Println(client)
	/*     srv, err := classroom.New(client)
	 *     if err != nil {
	 *         log.Fatalf("Unable to create classroom Client %v", err)
	 *     }
	 *
	 *     r, err := srv.Courses.List().PageSize(10).Do()
	 *     if err != nil {
	 *         log.Fatalf("Unable to retrieve courses. %v", err)
	 *     }
	 *     if len(r.Courses) > 0 {
	 *         fmt.Print("Courses:\n")
	 *         for _, c := range r.Courses {
	 *             fmt.Printf("%s (%s)\n", c.Name, c.Id)
	 *         }
	 *     } else {
	 *         fmt.Print("No courses found.")
	 *     } */
}
