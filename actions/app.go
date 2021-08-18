package actions

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"

	// i18n "github.com/gobuffalo/mw-i18n"
	paramlogger "github.com/gobuffalo/mw-paramlogger"

	l "github.com/tripleo1/repos/lib"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_repos_session",
		})
		// Automatically redirect to SSL
		// app.Use(forcessl.Middleware(secure.Options{
		// SSLRedirect:     ENV == "production",
		// SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		// }))

		if ENV == "development" {
			app.Use(paramlogger.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		// app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		// app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		// var err error
		// if T, err = i18n.New(packr.New("../locales", "../locales"), "en-US"); err != nil {
		// app.Stop(err)
		// }
		// app.Use(T.Middleware())

		app.GET("/", HomeHandler)

		// serve files from the public directory:

		app.GET("/clone", Clone)
		app.GET("/schedule", Schedule)
		app.GET("/jobs", Jobs)
		// app.Middleware.Skip(//HomeHandler,
		// Clone, Schedule, Jobs)

		// app.Resource("/items", ItemsResource{})

		app.ServeFiles("/", assetsBox)
	}

	return app
}

func Clone(c buffalo.Context) error {
	q := c.Param("q")
	d := c.Param("d")

	var root string
	if ENV == "production" {
		root = "/home/repos/Repos"
		os.Mkdir("/home/repos/Jobs", 0744)
	} else {
		root = "Repos"
		os.Mkdir("Jobs", 0744)
	}

	url1 := q
	depth := d == "1"

	bs, err := l.Get_bytes_for_url(root, url1, depth)
	if err != nil {
		//log.Fatal("Who knows")
		//return c.Render(500, r.HTML(...))
		return c.Render(500, r.String(err.Error()))
	}

	jn := l.Get_job_number()
	ioutil.WriteFile(path.Join("Jobs", jn), bs, 0600)

	return c.Render(200, r.String( /*"<pre>"+*/ string(bs) /*+"</pre>"*/)) //HTML("auth/new.plush.html"))
}

func Schedule(c buffalo.Context) error {
	// c.Set("user", models.User{})
	return c.Render(200, r.HTML("schedule.plush.html"))
}

func Jobs(c buffalo.Context) error {
	// c.Set("user", models.User{})
	return c.Render(200, r.HTML("jobs.plush.html"))
}
