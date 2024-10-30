package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/rizbo-dev/social-api/internal/store"
)

var usernames = []string{
	"Nikola", "Stefan", "Miloš", "Marko", "Jovan",
	"Luka", "Filip", "Nemanja", "Aleksandar", "Vuk",
	"Ana", "Jelena", "Milica", "Ivana", "Marija",
	"Maja", "Dragana", "Sofija", "Teodora", "Tamara",
	"Petar", "Mihajlo", "Bogdan", "Andrija", "Đorđe",
	"Iva", "Katarina", "Lidija", "Olga", "Nina",
}

var titles = []string{
	"Kako Početi Sa Programiranjem u Golangu",
	"Zašto Je Go Savršen Jezik za Backend Razvoj",
	"Uvod u Goroutine: Paralelizam u Golangu",
	"Kreiranje REST API-ja u Golangu",
	"10 Najboljih Paketa u Golangu za Početnike",
	"Rad sa JSON Podacima u Golangu",
	"Razumevanje Interfaces u Golangu",
	"Go Modules: Upravljanje Zavisnostima u Golangu",
	"Rad sa Datotekama u Golangu: Osnovni Saveti",
	"Razlika Između Struct i Map u Golangu",
	"Kako Kreirati Web Server u Golangu",
	"Error Handling u Golangu: Najbolje Prakse",
	"Primeri Rad sa Bazom Podataka u Golangu",
	"Rad sa CSV Fajlovima u Golangu",
	"Unit Testiranje u Golangu: Osnove",
	"HTTP Klijent u Golangu: Kako Poslati Zahtev",
	"Kako Efikasno Koristiti Goroutine i Kanale",
	"Razumevanje Struct Tagova u Golangu",
	"Šta Je Go Playground i Kako Ga Koristiti",
	"Kako Kreirati CLI Aplikaciju u Golangu",
	"Koraci za Implementaciju Docker-a sa Golang Aplikacijama",
	"Migracija Projekta na Go 1.18",
	"Korišćenje Go Gin Framework-a za Web Aplikacije",
	"Kako Se Pravilno Rukovati Greškama u Golangu",
	"Uvod u Cobra: Biblioteka za CLI Aplikacije u Golangu",
	"Poređenje Golanga i Node.js za Web Razvoj",
	"Razumevanje Pointers u Golangu",
	"Rad sa XML Podacima u Golangu",
	"Kako Postaviti Grafički Interfejs u Golangu",
	"Kako Optimizovati Golang Kod za Performanse",
}

var contents = []string{
	"Osnove Golanga: vodič kroz instalaciju, konfiguraciju i pisanje prvog programa.",
	"Prednosti korišćenja Golanga za backend razvoj i poređenje sa drugim jezicima.",
	"Kako koristiti Goroutines za upravljanje paralelizmom i performansama u aplikacijama.",
	"Kreiranje jednostavnog REST API-ja sa Golangom, uključujući osnovne HTTP metode.",
	"Top 10 paketa za produktivnost u Golangu koje svaki početnik treba da zna.",
	"Kako dekodirati i enkodirati JSON podatke u Golangu za rad sa API-ima.",
	"Vodič kroz interfaces u Golangu: koncepti, koristi i primeri.",
	"Upravljanje zavisnostima u Golangu sa Go Modules: najbolje prakse i primeri.",
	"Čitanje i pisanje datoteka u Golangu: korisni saveti i trikovi.",
	"Kada koristiti struct, a kada map u Golangu: primeri i poređenja.",
	"Kako kreirati jednostavan web server u Golangu i obraditi osnovne rute.",
	"Efikasno rukovanje greškama u Golangu: primeri i najbolje prakse.",
	"Rad sa bazom podataka u Golangu korišćenjem GORM paketa.",
	"Čitanje i pisanje CSV fajlova u Golangu: primeri i praktična primena.",
	"Vodič kroz unit testiranje u Golangu: kako testirati funkcije i module.",
	"Kako poslati HTTP zahtev u Golangu i obraditi odgovore sa servera.",
	"Rad sa kanalom (channels) u Golangu za komunikaciju između gorutina.",
	"Razumevanje struct tagova u Golangu: njihova primena i koristi.",
	"Kako koristiti Go Playground za testiranje i deljenje koda.",
	"Kreiranje jednostavne CLI aplikacije u Golangu koristeći Cobra paket.",
	"Integracija Docker-a sa Golang aplikacijama za lakše razmeštanje.",
	"Nove funkcionalnosti i poboljšanja u Go 1.18: šta možete očekivati.",
	"Korišćenje Gin framework-a za kreiranje API-ja u Golangu: osnovni primer.",
	"Pravilno rukovanje greškama u Golangu za stabilnije aplikacije.",
	"Kako Cobra biblioteka pomaže u razvoju moćnih CLI alata u Golangu.",
	"Poređenje performansi Golanga i Node.js za razvoj backend servisa.",
	"Rad sa pokazivačima (pointers) u Golangu: koncepti i praktični primeri.",
	"Parsiranje i generisanje XML podataka u Golangu za razmenu informacija.",
	"Vodič kroz postavljanje GUI aplikacija u Golangu koristeći Fyne biblioteku.",
	"Kako optimizovati Golang kod za bolje performanse: korisni saveti.",
}

var tags = []string{
	"Golang", "Programiranje", "Backend", "API", "Web Razvoj",
	"Goroutines", "Paralelizam", "Performanse", "JSON", "Interfaces",
	"Go Modules", "Zavisnosti", "Datoteke", "Struct", "Mape",
	"Web Server", "Greške", "Baze Podataka", "CSV", "Testiranje",
	"HTTP Zahtevi", "Kanali", "Struct Tagovi", "Go Playground", "CLI Aplikacije",
	"Docker", "Gin Framework", "Razvoj Softvera", "Pokazivači", "Optimizacija",
}

var commentsData = []string{
	"Odličan članak! Puno mi je pomogao da razumem osnove Golanga.",
	"Hvala na detaljnom objašnjenju gorutina! Konačno mi je sve jasnije.",
	"Super primeri, posebno za rad sa JSON podacima. Bravo!",
	"Da li možeš da napišeš više o unit testiranju u Golangu?",
	"Ovaj vodič za REST API je fantastičan! Svaka čast!",
	"Interesuje me kako koristiš Docker sa Golang aplikacijama.",
	"Veoma koristan članak o greškama u Golangu, hvala puno!",
	"Da li ćeš raditi tutorial o kanalu u Golangu? Zvuči interesantno!",
	"Hvala na savetima za optimizaciju, definitivno ću ih isprobati.",
	"Prvi put razumem strukture i mape. Sjajno objašnjeno!",
	"Da li postoji način da dublje objasniš rad sa Go modules?",
	"Bilo bi sjajno kada bi napisao još o radu sa bazama podataka.",
	"Kratko i jasno objašnjeno, svaka čast! Puno mi je pomoglo.",
	"Bravo za odličan članak o HTTP zahtevima u Golangu!",
	"Interesuje me više o razlici između Gorutina i običnih funkcija.",
	"Kako mogu da implementiram ovaj primer u svom projektu?",
	"Hvala što si podelio iskustvo sa greškama! Veoma korisno.",
	"Članak o pokazivačima mi je baš pomogao! Možeš li dodatno objasniti?",
	"Da li imaš savet za rad sa velikim JSON fajlovima?",
	"Mnogo korisnih informacija u vezi sa Dockerom, hvala!",
	"Rad sa CSV-om mi je zadao glavobolje, ali sada je sve jasno.",
	"Voleo bih da pročitam više o Cobra frameworku.",
	"Sjajno objašnjenje, baš sam naučio mnogo novih stvari o Go Playgroundu!",
	"Primeri su odlični, ali da li imaš više saveta za testiranje?",
	"Kako koristiti ovaj kod za rad sa Gin frameworkom?",
	"Saveti za optimizaciju performansi su odlično objašnjeni.",
	"Korisne informacije o razvoju CLI aplikacija. Može li još primera?",
	"Voleo bih da vidim više o paralelnom radu sa gorutinama.",
	"Hvala na postu o struct tagovima, veoma je korisno za rad sa bazama.",
	"Preporučujem ovaj blog svakome ko želi da nauči Golang!",
}

func Seed(store store.Storage) error {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
		}
	}

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
		}
	}

	log.Println("Seeding complete")

	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.IntN(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.IntN(len(titles))],
			Content: contents[rand.IntN(len(contents))],
			Tags: []string{
				tags[rand.IntN(len(tags))],
				tags[rand.IntN(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		comments[i] = &store.Comment{
			PostID:  posts[rand.IntN(len(posts))].ID,
			UserID:  users[rand.IntN(len(users))].ID,
			Content: commentsData[rand.IntN(len(commentsData))],
		}
	}

	return comments
}
