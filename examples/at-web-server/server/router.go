package server

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/admin/team"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/content"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/dashboard"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/device"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/login"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/room"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/user"
	"github.com/go-chi/chi"
)

func (s *Server) NewRouter() (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(middleware.LoggerContext(s.opts.logger))
	mux.Use(middleware.GRPCConnContext(s.conn))

	// root "/"
	mux.Method(http.MethodGet, "/", http.RedirectHandler(s.BuildURL("/dashboard").String(), http.StatusSeeOther))

	// css & js
	mux.Mount("/static/", http.FileServer(http.FS(StaticFS())))

	// without login
	mux.Get("/login", login.Handler(s.config.Firebase))
	mux.Post("/login", auth.LoginHandler(s.client, s.config.EnableTLS))
	mux.Get("/logout", auth.LogoutHandler(s.client, s))

	// require login
	mux.Group(func(r chi.Router) {
		r.Use(auth.UserProfileContext(s.client, s))

		r.Get("/dashboard", dashboard.Handler(s))

		// settings
		r.Route("/settings", func(r chi.Router) {
			// user settings
			r.Route("/user", func(r chi.Router) {
				r.Get("/", user.Handler())

				r.Get("/edit", user.EditPageHandler())
				r.Post("/edit", user.EditHandler(s.client, s))

				r.Get("/password", user.ChangePasswordPageHandler())
				r.Post("/password", user.ChangePasswordHandler(s.client, s))

				r.Get("/device", user.SetDevicePageHandler())
				r.Post("/device", user.SetDeviceHandler(s.client, s))
			})

			// room settings
			r.Route("/rooms", func(r chi.Router) {
				r.Get("/", room.ListPageHandler())
				r.Get("/new", room.CreatePageHandler())
				r.Post("/new", room.CreateHandler(s))

				r.Route("/{roomID}", func(r chi.Router) {
					r.Use(middleware.RoomContext)

					r.Get("/", room.EditPageHandler())
					r.Post("/", room.EditHandler(s))
					r.Get("/delete", room.DeletePageHandler())
					r.Post("/delete", room.DeleteHandler(s))
					r.Get("/devices", room.SetDevicePageHandler())
					r.Post("/devices", room.SetDeviceHandler(s))
				})
			})

			r.Route("/devices", func(r chi.Router) {
				// device settings
				r.Get("/", device.ListPageHandler())
				r.Get("/new", device.CreatePageHandler())
				r.Post("/new", device.CreateHandler(s))

				r.Route("/{deviceID}", func(r chi.Router) {
					r.Use(middleware.DeviceContext)

					r.Get("/", device.EditPageHandler())
					r.Post("/", device.EditHandler(s))
					r.Get("/delete", device.DeletePageHandler())
					r.Post("/delete", device.DeleteHandler(s))
				})
			})

			r.Route("/contents", func(r chi.Router) {
				// content settings
				r.Get("/", content.ListPageHandler())
				r.Get("/new", content.CreatePageHandler())
				r.Post("/new", content.CreateHandler(s))

				r.Route("/{contentID}", func(r chi.Router) {
					r.Use(middleware.ContentContext)

					r.Get("/", content.EditPageHandler())
					r.Post("/", content.EditHandler(s))
					r.Get("/delete", content.DeletePageHandler())
					r.Post("/delete", content.DeleteHandler(s))
				})
			})
		})

		// for admin
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.AdminContext)

			r.Route("/teams", func(r chi.Router) {
				r.Get("/", team.ListPageHandler())
				r.Get("/new", team.CreatePageHandler())
				r.Post("/new", team.CreateHandler(s))

				r.Route("/{teamID}", func(r chi.Router) {
					r.Use(middleware.TeamContext())

					r.Get("/", team.EditPageHandler())
					r.Post("/", team.EditHandler(s))
					r.Get("/delete", team.DeletePageHandler())
					r.Post("/delete", team.DeleteHandler(s))
				})
			})
		})

		r.Route("/api", func(r chi.Router) {
			r.Route("/dashboard", func(r chi.Router) {
				r.Get("/", dashboard.DashboardAPIHandler())
				r.Get("/event", dashboard.EventStreamHandler())
				r.Post("/stop", dashboard.StopAPIHandler())

				r.Route("/rooms", func(r chi.Router) {
					r.Route("/{roomID}", func(r chi.Router) {
						r.Use(middleware.RoomContext)

						r.Post("/join", dashboard.JoinAPIHandler())
						r.Post("/leave", dashboard.LeaveAPIHandler())
					})
				})

				r.Route("/contents", func(r chi.Router) {
					r.Route("/{contentID}", func(r chi.Router) {
						r.Use(middleware.ContentContext)

						r.Post("/play", dashboard.StopAPIHandler())
					})
				})
			})
		})
	})

	return mux, nil
}
