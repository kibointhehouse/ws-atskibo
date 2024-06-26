package url

import (
	"github.com/kibointhehouse/ws-atskibo/controller"

	"github.com/gofiber/fiber/v2"
)	
func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik

	//page.Get("/presensi", controller.GetPresensi)

	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
}
