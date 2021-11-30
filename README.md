# 👋 Welcome

# TERBILANG
Terbilang adalah package untuk mengubah nominal angka rupiah ke dalam bentuk teks 

<a href="https://pkg.go.dev/github.com/ekokurniadi/terbilang"><img src="https://pkg.go.dev/badge/github.com/ekokurniadi/terbilang.svg" alt="Go Reference"></a>

### How to install
```sh
go get -u github.com/ekokurniadi/terbilang
```

### Then import the package

```go
import "github.com/ekokurniadi/terbilang"
```
### Example Project

```go
package main

import (
	"fmt"
	"github.com/ekokurniadi/terbilang"
)

func main() {
    angka  := 50000
    format := terbilang.Init()
    fmt.Println(format.Convert(angka))
    
    //output => lima puluh ribu rupiah
}

```
### Run example project and you will see the output
```sh
go run main.go
lima puluh ribu rupiah
```

<<<<<<< HEAD
### Have Fun :D

=======
### Example coding when i implement terbilang on my project

```go
package campaign

import (
	"strings"
        "github.com/ekokurniadi/terbilang"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageUrl         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
        GAmountTerbilang string                    `json:"amount_on_text`
	CurrentAmount    int                       `json:"current_amount"`
	BackerCount      int                       `json:"backer_count"`
	UserID           int                       `json:"user_id"`
	Slug             string                    `json:"slug"`
	User             CampaignUserFormatter     `json:"user"`
	Perk             []string                  `json:"perks"`
	Images           []CampaignImagesFormatter `json:"images"`
}
type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.BackerCount = campaign.BakerCount
        
        //this coding how to use terbilang package
        //t.Generate() is a function with callback a string
        t := terbilang.Init()
        campaignDetailFormatter.GAmount = t.Convert(int64(campaign.GoalAmount))
	
        campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perk = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImagesFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImagesFormatter := CampaignImagesFormatter{}
		campaignImagesFormatter.ImageUrl = image.FileName
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImagesFormatter.IsPrimary = isPrimary

		images = append(images, campaignImagesFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter

}
```

### Example response when i using terbilang package


```json 
{
    "meta": {
        "message": "Campaign Detail",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "name": "Campaign 1",
        "short_description": "short",
        "description": "description",
        "image_url": "satu.jpg",
        "goal_amount": 100000,
        "amount_on_text": "seratus ribu rupiah",
        "current_amount": 0,
        "backer_count": 0,
        "user_id": 12,
        "slug": "campaign-1",
        "user": {
            "name": "Eko Kurniadi",
            "image_url": "images/1-laporan.png"
        },
        "perks": [
            "satu",
            "dua",
            "tiga"
        ],
        "images": [
            {
                "image_url": "satu.jpg",
                "is_primary": false
            },
            {
                "image_url": "images/1-transaksi.PNG",
                "is_primary": false
            },
            {
                "image_url": "images/1-dashboard.PNG",
                "is_primary": false
            },
            {
                "image_url": "images/1-dashboard.PNG",
                "is_primary": true
            },
            {
                "image_url": "images/1-dashboard.PNG",
                "is_primary": false
            }
        ]
    }
}
```
>>>>>>> 751a311fdad7e02687fdc22131bf213e9db15ab5
