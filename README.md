## storemeta


- [x] This is a simple golang script that will parse/scrape the apps-store/play-store  meta-infos:

####        a. Android Play Store meta info

######      https://play.google.com/store/apps/details?id=com.google.android.apps.photos&hl=en


####        b. Apple App Store meta info 

######      https://itunes.apple.com/app/id293622097?mt=8

- [x] Output the meta-info in JSON format



## Compile

```sh

     git clone https://github.com/bayugyug/storemeta.git && cd storemeta 

     git pull && make 

```


## Meta-Infos

```go

        type App struct {
                Platform        string `json:"platform"`
                AppURL          string `json:"app-url"`
                AppID           string `json:"app-id"`
                Genre           string `json:"genre"`
                Title           string `json:"title"`
                Description     string `json:"description"`
                Badge           string `json:"badge"`
                RatingTotal     string `json:"rating-total"`
                RatingPerStar   string `json:"rating-per-star"`
                RatingDesc      string `json:"rating-desc"`
                RatingValue     string `json:"rating-value"`
                SoftwarePrice   string `json:"software-price"`
                MetaDesc        string `json:"meta-desc"`
                FileSize        string `json:"file-size"`
                ContentRating   string `json:"content-rating"`
                DatePublished   string `json:"date-published"`
                SoftwareVersion string `json:"software-version"`
                SoftwareOs      string `json:"software-os"`
                TotalDownloads  string `json:"total-downloads"`
                Developer       string `json:"developer"`
                DeveloperSite   string `json:"developer-site"`
        }

```





## Usage

```sh

$ ./storemeta -h

Ver: 0.1.0-20160830.022006

  -a string
        use for querying the Android App Store (shorthand)
  -android string
        use for querying the Android App Store
  -ca string
        use for querying the list of apps per category in IOS App Store (shorthand)
  -category-android string
        use for querying the list of apps per category in IOS App Store
  -category-ios string
        use for querying the list of apps per category in IOS App Store
  -ci string
        use for querying the list of apps per category in IOS App Store (shorthand)
  -d    use to enable the output in console (shorthand) (default true)
  -debug
        use to enable the output in console (default true)
  -h    Show this help/how-to
  -i string
        use for querying the IOS App Store (shorthand)
  -ios string
        use for querying the IOS App Store
  -la
        use for querying the category list of apps in Android App Store (shorthand)
  -li
        use for querying the category list of apps in IOS App Store (shorthand)
  -list-category-android
        use for querying the category list of apps in Android App Store
  -list-category-ios
        use for querying the category list of apps in IOS App Store
  -pf string
        use to enable what format is used in showing the output (shorthand) (default "json")
  -print-format string
        use to enable what format is used in showing the output (default "json")


        Example:

                ./storemeta -a <AndroidStoreID>  -i <IOSStoreID>

                ./storemeta -list-category-android

                ./storemeta -list-category-ios

                ./storemeta -category-android=GAME_ACTION

                ./storemeta -category-ios=GAMES_ACTION

                ./storemeta  -a="com.google.android.apps.photos"

                or

                ./storemeta  -i="293622097"

                or

                ./storemeta  -a="com.google.android.apps.photos" -i="293622097"

                or

                ./storemeta  -a="com.google.android.apps.plus,com.google.android.launcher,com.sphero.sprk"

                or

                ./storemeta  -i="544007664,535886823,643496868"

```


## Sample (Android)



#### Simple
```go


$    ./storemeta  -a="com.google.android.apps.photos"
    {
        "platform": "ANDROID",
        "title": "Google Photos - Android Apps on Google Play",
        "developer": "Google Inc.",
        "developer-site": "https://www.google.com/url?q=http://www.google.com/policies/privacy\u0026sa=D\u0026usg=AFQjCNE7y6nm7TcHvct7CDJRmWrYBHvMEQ",
        "genre": "PHOTOGRAPHY",
        "description": "Google Photos is the home for all your photos and videos. Automatically organized and searchable, you can find photos fast and bring them to life. It’s the photo gallery that thinks like you do.  VISUAL SEARCH Your photos are now searchable by the places and things that appear in them. Looking for that fish taco you ate in Hawaii? Just search “food in Hawaii” to find it – no tagging required.  UNLIMITED FREE HIGH QUALITY STORAGE Automatically backup all your photos and videos. Access them on any device or on the web at photos.google.com. Your photos are safe, secure, and private to you.  FREE UP SPACE ON YOUR DEVICE Never worry about running out of space on your phone again. In Settings, just tap “Free up device storage” – photos that are safely backed up will be removed from your device’s storage, but will still be available in Google Photos.  BRING PHOTOS TO LIFE Enjoy automatically created montage movies, interactive stories, collages, animations, panoramas, and more from your photos. Or you can easily create them yourself – just tap +.  EASY EDITING Transform photos with the tap of a finger. Use simple, yet powerful, photo and video editing tools to apply filters, adjust colors, and more.  SHARED ALBUMS Get everyone’s photos and videos in one place, across Android, iOS, and the web. Privately sharing all the photos you took – and getting the ones you didn’t – has never been easier.  INSTANT SHARING Instantly share up to 1,500 photos with anyone – no matter what device they’re on. In the share menu, just tap Create Link.  REDISCOVER YOUR PHOTOS It’s easier than ever to relive your memories. The Assistant can create collages of your old photos that help you relive the past.  READY TO CAST View your photos and videos on your TV with Chromecast support.  Follow us for the latest news and updates Twitter: https://twitter.com/GooglePhotos Google+: https://google.com/+GooglePhotos  Need help? Visit https://support.google.com/photos  Face grouping is not available in all countries.",
        "badge": "Top Developer",
        "rating-total": "4846787",
        "rating-per-star": "3214313,920228,344084,130721,237441",
        "rating-desc": "Rated 4.4 stars out of five stars\nRated 5 stars out of five stars\nRated 4 stars out of five stars\nRated 2 stars out of five stars\nRated 3 stars out of five stars\nRated 1 stars out of five stars\nRated 4.3 stars out of five stars\nRated 4.5 stars out of five stars\nRated 4.2 stars out of five stars\nRated 4.6 stars out of five stars\nRated 3.9 stars out of five stars\nRated 4.1 stars out of five stars\nRated 4.0 stars out of five stars",
        "rating-value": "4.391282558441162",
        "software-price": "0",
        "meta-desc": "\r\n Updated August 22, 2016 \r\n Installs   500,000,000 - 1,000,000,000   \r\n Requires Android       4.0 and up     \r\n Content Rating Rated for 3+  Learn more  \r\n Permissions  View details  \r\n Report  Flag as inappropriate  \r\n  Offered By  Google Inc. \r\n  Developer    Visit website   Email apps-help@google.com   Privacy Policy  1600 Amphitheatre Parkway, Mountain View 94043  ",
        "file-size": "",
        "content-rating": "Rated for 3+",
        "date-published": "",
        "software-version": "",
        "software-os": "4.0 and up",
        "total-downloads": "750000000",
        "app-url": "https://play.google.com/store/apps/details?id=com.google.android.apps.photos\u0026hl=en",
        "app-id": "com.google.android.apps.photos"
   }
```

## Sample (IOS)

#### Simple
```go


$   ./storemeta  -i="293622097"

    {
        "platform": "IOS",
        "title": "Google Earth on the App StoreGoogle Earth\n              By Google, Inc.\n              \n            \n            \n              \n              View More by This Developer\n              \n            \n            Open iTunes to buy and download apps.",
        "developer": "Google, Inc.",
        "developer-site": "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewEula?id=293622097",
        "genre": "TRAVEL",
        "description": "Fly around the planet with a swipe of your finger with Google Earth for iPhone, iPad, and iPod touch. Explore distant lands or reacquaint yourself with your childhood home. Search for cities, places, and businesses. Browse layers including roads, borders, places, photos and more.  See the world at street level with integrated Street View.Use the “tour guide” to easily discover exciting new places to explore. With a quick swipe on the tab at the bottom of the screen, you can bring up a selection of virtual tours from around the globe.With 3D imagery, you can now fly through complete 3D recreations of select cities, including San Francisco, Boston, Rome, and others. With every building modeled in 3D, you truly get a sense of flying above the city.",
        "badge": "This app is designed for both iPhone and iPad",
        "rating-total": "446047",
        "rating-per-star": "378,445669",
        "rating-desc": "3 and a half stars, 378 Ratings",
        "rating-value": "3.69312",
        "software-price": "Free",
        "meta-desc": "",
        "file-size": "30.1",
        "content-rating": "Rated 4+",
        "date-published": "2016-05-03 00:00:00",
        "software-version": "7.1.6",
        "software-os": "Requires iOS 5.0 or later. Compatible with iPhone, iPad, and iPod touch.",
        "total-downloads": "44604700",
        "app-url": "https://itunes.apple.com/us/app/google-earth/id293622097?mt=8",
        "app-id": "293622097"
    }
```


```

## List the Apps Categories

#### (IOS)

```go


$   ./storemeta -print-format=simple -list-category-ios
        1 BOOKS
        2 BUSINESS
        3 EDUCATION
        4 ENTERTAINMENT
        5 FINANCE
        6 FOOD_DRINK
        7 GAMES
        8 GAMES_ACTION
        9 GAMES_ADVENTURE
        10 GAMES_ARCADE
        11 GAMES_BOARD
        12 GAMES_CARD
        13 GAMES_CASINO
        14 GAMES_DICE
        15 GAMES_EDUCATIONAL
        16 GAMES_FAMILY
        17 GAMES_MUSIC
        18 GAMES_PUZZLE
        19 GAMES_RACING
        20 GAMES_ROLE_PLAYING
        21 GAMES_SIMULATION
        22 GAMES_SPORTS
        23 GAMES_STRATEGY
        24 GAMES_TRIVIA
        25 GAMES_WORD
        26 HEALTH_FITNESS
        27 LIFESTYLE
        28 MAGAZINES_NEWSPAPERS
        29 MEDICAL
        30 MUSIC
        31 NAVIGATION
        32 NEWS
        33 PHOTO_VIDEO
        34 PRODUCTIVITY
        35 REFERENCE
        36 SHOPPING
        37 SOCIAL_NETWORKING
        38 SPORTS
        39 TRAVEL
        40 UTILITIES
        41 WEATHER

```


#### (ANDROID)

```go

$       ./storemeta  -print-format=simple -list-category-android

        1 ANDROID_WEAR
        2 BOOKS_AND_REFERENCE
        3 BUSINESS
        4 COMICS
        5 COMMUNICATION
        6 EDUCATION
        7 ENTERTAINMENT
        8 FINANCE
        9 HEALTH_AND_FITNESS
        10 LIBRARIES_AND_DEMO
        11 LIFESTYLE
        12 MEDIA_AND_VIDEO
        13 MEDICAL
        14 MUSIC_AND_AUDIO
        15 NEWS_AND_MAGAZINES
        16 PERSONALIZATION
        17 PHOTOGRAPHY
        18 PRODUCTIVITY
        19 SHOPPING
        20 SOCIAL
        21 SPORTS
        22 TOOLS
        23 TRANSPORTATION
        24 TRAVEL_AND_LOCAL
        25 WEATHER
        26 GAME
        27 GAME_ACTION
        28 GAME_ADVENTURE
        29 GAME_ARCADE
        30 GAME_BOARD
        31 GAME_CARD
        32 GAME_CASINO
        33 GAME_CASUAL
        34 GAME_EDUCATIONAL
        35 GAME_MUSIC
        36 GAME_PUZZLE
        37 GAME_RACING
        38 GAME_ROLE_PLAYING
        39 GAME_SIMULATION
        40 GAME_SPORTS
        41 GAME_STRATEGY
        42 GAME_TRIVIA
        43 GAME_WORD
        44 FAMILY
        45 FAMILY_AGE_AGE_RANGE1
        46 FAMILY_AGE_AGE_RANGE2
        47 FAMILY_AGE_AGE_RANGE3
        48 FAMILY_ACTION
        49 FAMILY_BRAINGAMES
        50 FAMILY_CREATE
        51 FAMILY_EDUCATION
        52 FAMILY_MUSICVIDEO
        53 FAMILY_PRETEND



```


## List the App in the Category

#### (IOS)

```go

$ ./storemeta -category-ios=GAMES_ACTION
[        {
                "preview": "https://itunes.apple.com/us/app/call-of-mini-zombies/id431213733?mt=8",
                "store_id": "431213733",
                "category": "GAMES_ACTION"
        },
        {
                "preview": "https://itunes.apple.com/us/app/gun-club-2-best-in-virtual/id311594640?mt=8",
                "store_id": "311594640",
                "category": "GAMES_ACTION"
        },
        {
                "preview": "https://itunes.apple.com/us/app/ace-fishing-wild-catch/id694972182?mt=8",
                "store_id": "694972182",
                "category": "GAMES_ACTION"
        },
        {
                "preview": "https://itunes.apple.com/us/app/respawnables/id575684686?mt=8",
                "store_id": "575684686",
                "category": "GAMES_ACTION"
        },
        {
                "preview": "https://itunes.apple.com/us/app/swing/id1064078609?mt=8",
                "store_id": "1064078609",
                "category": "GAMES_ACTION"
        }
]

```

#### (ANDROID)

```go
$ ./storemeta -category-android=FAMILY_AGE_AGE_RANGE1
[
        {
                "preview": "https://play.google.com/store/apps/details?id=com.disney.PaintAndPlay_goo",
                "store_id": "com.disney.PaintAndPlay_goo",
                "category": "FAMILY_AGE_AGE_RANGE1"
        },
        {
                "preview": "https://play.google.com/store/apps/details?id=com.intellijoy.android.shapes",
                "store_id": "com.intellijoy.android.shapes",
                "category": "FAMILY_AGE_AGE_RANGE1"
        },
        {
                "preview": "https://play.google.com/store/apps/details?id=com.intellijoy.android.shapes",
                "store_id": "com.intellijoy.android.shapes",
                "category": "FAMILY_AGE_AGE_RANGE1"
        },
        {
                "preview": "https://play.google.com/store/apps/details?id=com.intellijoy.android.shapes",
                "store_id": "com.intellijoy.android.shapes",
                "category": "FAMILY_AGE_AGE_RANGE1"
        },
        {
                "preview": "https://play.google.com/store/apps/details?id=com.intellijoy.android.shapes",
                "store_id": "com.intellijoy.android.shapes",
                "category": "FAMILY_AGE_AGE_RANGE1"
        }
]



```

## Note:

- [x] IOS Total Downloads is calculated :

        Rating-Total x 10

- [x] Android Total Downloads is calculated :
    
        ( As Is in "Installs" )


## Docker Binary

- [x] In order to  use it as dockerize binary


```sh

    sudo  sysctl -w net.ipv4.ip_forward=1

    sudo  docker run --rm  registry.hub.docker.com/bayugyug/storemeta -h

    sudo  docker run --rm  registry.gitlab.com/bayugyug/storemeta -h

```



## As HTTP Server

- [x] In order to  use it via CURL/WGET or Browser


```sh

    sudo  sysctl -w net.ipv4.ip_forward=1

    sudo  docker run -p 7000-8000:7000-8000 -v `pwd`:`pwd` -w `pwd` -d --name storemeta-latest  registry.hub.docker.com/bayugyug/storemeta:latest --http --port 7778

    sudo  docker run -p 7000-8000:7000-8000 -v `pwd`:`pwd` -w `pwd` -d --name storemeta-latest  registry.gitlab.com/bayugyug/storemeta:latest --http --port 7778

    curl -i -v 'http://127.0.0.1:7778/list-category-android'

    curl -i -v 'http://127.0.0.1:7778/list-category-ios'

    curl -i -v 'http://127.0.0.1:7778/category-android/?p=GAME_ACTION'

    curl -i -v 'http://127.0.0.1:7778/category-ios/?p=GAMES_ACTION'

    curl -i -v 'http://127.0.0.1:7778/storeid/?a=com.google.android.apps.photos&i=293622097'

```

### License

[MIT](https://bayugyug.mit-license.org/)

