package basics

type ScrapedData struct {
	Content   string
	TimeTaken int
}

type Scrapper interface {
	Scrape(source string) ScrapedData
}

type WebScrapper struct {
}

func (scrapper WebScrapper) Scrape(source string) ScrapedData {
	var scraped_data ScrapedData
	scraped_data.Content = "<html>asas</html>"
	scraped_data.TimeTaken = 1
	return scraped_data
}

type DriveScrapper struct {
}

func (scrapper DriveScrapper) Scrape(source string) ScrapedData {
	return ScrapedData{
		Content:   "All data",
		TimeTaken: 1,
	}
}

var ScrapperFactory = map[string]Scrapper{
	"web":   WebScrapper{},
	"drive": DriveScrapper{},
}
