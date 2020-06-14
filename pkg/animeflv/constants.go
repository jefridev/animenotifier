package animeflv

const hostname = "https://animeflv.net"

// List of xpath used to extract anime information.
const (
	endpointXpath      = "//ul[@class='ListSdbr']//li//a"
	titleXpath         = "//h1[@class='Title']"
	typeXpath          = "//span[contains(@class,'Type')]"
	coverXpath         = "//div[@class='AnimeCover']//figure//img/@src"
	statusXpath        = "//p[@class='AnmStts']/span/text()"
	synopsisXpath      = "//div[@class='Description']/p/text()"
	genresXpath        = "//nav[@class='Nvgnrs']//a"
	scriptContentXpath = "/html/body/script[12]"
)
