## storemeta


- [x] This is a simple golang script that will parse/scrape the apps-store/play-store  meta-infos:

####        a. Android Play Store meta info

######      https://play.google.com/store/apps/details?id=com.google.android.apps.photos&hl=en


####        b. Apple App Store meta info 

######      https://itunes.apple.com/app/id293622097?mt=8

- [x] Output the meta-info in JSON format


## Note:

- [x] IOS Total Downloads is calculated :

        Rating-Total x 10

- [x] Android Total Downloads is calculated :
        Average of Numdownloads which is 
    
        ( NumDownloads From + NumDownloads To )  / 2



## Compile

```sh

     git clone https://github.com/bayugyug/storemeta.git && cd storemeta 

     git pull && make 

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


#### Multiple
```go

$       ./storemeta  -a="com.google.android.apps.plus,com.google.android.launcher,com.sphero.sprk"
{
        "platform": "ANDROID",
        "title": "Google+ - Android Apps on Google Play",
        "developer": "Google Inc.",
        "developer-site": "https://www.google.com/url?q=http://www.google.com/policies/privacy\u0026sa=D\u0026usg=AFQjCNE7y6nm7TcHvct7CDJRmWrYBHvMEQ",
        "genre": "SOCIAL",
        "description": "Discover amazing things created by passionate people.  • Explore your interests • Join Communities of people around any topic • Group things you love into Collections • Build a home stream filled with amazing content",
        "badge": "Top Developer",
        "rating-total": "2744541",
        "rating-per-star": "1698276,467070,234151,94514,250523",
        "rating-desc": "Rated 4.2 stars out of five stars\nRated 4 stars out of five stars\nRated 2 stars out of five stars\nRated 1 stars out of five stars\nRated 3 stars out of five stars\nRated 5 stars out of five stars\nRated 4.0 stars out of five stars\nRated 4.5 stars out of five stars\nRated 4.3 stars out of five stars\nRated 3.9 stars out of five stars\nRated 4.4 stars out of five stars\nRated 3.7 stars out of five stars\nRated 4.1 stars out of five stars",
        "rating-value": "4.190752983093262",
        "software-price": "0",
        "meta-desc": "\r\n Updated August 19, 2016 \r\n Size    Varies with device   \r\n Installs   1,000,000,000 - 5,000,000,000   \r\n Current Version    Varies with device   \r\n Requires Android    Varies with device     \r\n Content Rating Rated for 12+ Parental Guidance Recommended Learn more  \r\n  Interactive Elements  Users Interact, Shares Location \r\n Permissions  View details  \r\n Report  Flag as inappropriate  \r\n  Offered By  Google Inc. \r\n  Developer    Visit website   Email apps-help@google.com   Privacy Policy  1600 Amphitheatre Parkway, Mountain View 94043  ",
        "file-size": "",
        "content-rating": "Rated for 12+",
        "date-published": "",
        "software-version": "Varies with device",
        "software-os": "Varies with device",
        "total-downloads": "3000000000",
        "app-url": "https://play.google.com/store/apps/details?id=com.google.android.apps.plus\u0026hl=en",
        "app-id": "com.google.android.apps.plus"
}
{
        "platform": "ANDROID",
        "title": "SPRK Lightning Lab for Sphero - Android Apps on Google Play",
        "developer": "Sphero",
        "developer-site": "https://www.google.com/url?q=http://sphero.com/privacy\u0026sa=D\u0026usg=AFQjCNEZMVaQnl-ALYvWQt4GzazDqrGPuw",
        "genre": "EDUCATION",
        "description": "This is SPRK Lightning Lab. Your hub to create, contribute, and learn with Sphero robots.Simple for beginners yet sophisticated enough for seasoned programmers, Lightning Lab empowers anyone to program their robot. The visual block-based building app makes learning the basic principles of programming approachable and fun. Drag and drop actions, controls, operators, and more to learn how your Sphero works. The written code you create can then be viewed right beside your block sequence for a more advanced understanding of our C-based language. Join the growing community of makers, students, instructors, and parents – all learning on the same social platform. Share your creations, comment on posted activities, and collaborate with other users around the globe to innovate the world of education. Join Lightning Lab and be a part of something bigger.Program your robotTransform ideas into code by using visual blocks that represent our C-based language, Oval. Double tap on a block to learn what it does. Complete awesome activitiesProgram a painting. Navigate a maze. Mimic the solar system. Swim across the water. Have a dance party… The only limit is your imagination. Take a driveNeed a brain break? Go Drive and play.“There were plenty of educational robots on the market, but Sphero stood out for the simplicity of its coding; moreover, unlike robotic cars, tanks, or dolls, a ball was equally inviting to boys and girls.” - The New Yorker“Simply put, kids could not only learn about programming, but also have fun doing so.” - Engadget Learning is evolving. Get on the ball. Order a SPRK+ robot today at sphero.com.*Supported Robots: SPRK+, SPRK Edition, Sphero 2.0, and Ollie**Supported Languages: English, French, Italian, German, Spanish, Chinese Simplified, Chinese Traditional, Japanese and Korean",
        "badge": "",
        "rating-total": "381",
        "rating-per-star": "239,61,15,15,51",
        "rating-desc": "Rated 4.1 stars out of five stars\nRated 4 stars out of five stars\nRated 5 stars out of five stars\nRated 1 stars out of five stars\nRated 2 stars out of five stars\nRated 3 stars out of five stars\nRated 3.2 stars out of five stars\nRated 3.5 stars out of five stars\nRated 4.2 stars out of five stars\nRated 2.9 stars out of five stars\nRated 2.2 stars out of five stars\nRated 4.3 stars out of five stars\nRated 3.1 stars out of five stars\nRated 4.4 stars out of five stars\nRated 3.9 stars out of five stars\nRated 4.5 stars out of five stars\nRated 3.4 stars out of five stars\nRated 3.7 stars out of five stars\nRated 3.8 stars out of five stars\nRated 3.0 stars out of five stars\nRated 3.3 stars out of five stars",
        "rating-value": "4.107611656188965",
        "software-price": "0",
        "meta-desc": "\r\n Updated July 26, 2016 \r\n Size  49M   \r\n Installs   10,000 - 50,000   \r\n Current Version  2.0.0   \r\n Requires Android       4.4 and up     \r\n Content Rating Rated for 3+  Learn more  \r\n  Interactive Elements  Users Interact \r\n Permissions  View details  \r\n Report  Flag as inappropriate  \r\n  Offered By  Sphero \r\n  Developer    Visit website   Email support@sphero.com   Privacy Policy  4772 Walnut St\nSte. 206\nBoulder, CO 80301  ",
        "file-size": "49",
        "content-rating": "Rated for 3+",
        "date-published": "",
        "software-version": "2.0.0",
        "software-os": "4.4 and up",
        "total-downloads": "30000",
        "app-url": "https://play.google.com/store/apps/details?id=com.sphero.sprk\u0026hl=en",
        "app-id": "com.sphero.sprk"
}
{
        "platform": "ANDROID",
        "title": "Google Now Launcher - Android Apps on Google Play",
        "developer": "Google Inc.",
        "developer-site": "https://www.google.com/url?q=http://www.google.com/policies/privacy\u0026sa=D\u0026usg=AFQjCNE7y6nm7TcHvct7CDJRmWrYBHvMEQ",
        "genre": "TOOLS",
        "description": "Upgrade the launcher on your Android device for a fast, clean home screen that puts Google Now just a swipe away.Available on all devices with Android 4.1 (Jelly Bean) or higher.Key features:• Swipe right from your Home screen to see Google Now cards that bring you just the right information, at just the right time.• Quick access to Search from every Home screen.• Say “Ok Google” to search with your voice, or tell your phone what to do: send a text message, get directions, play a song, and much more.• A-Z apps list, with fast scrolling and quick searching of apps on your device and the Play Store.• App Suggestions bring the app you’re looking for to the top of your A-Z list.",
        "badge": "Top Developer",
        "rating-total": "706588",
        "rating-per-star": "451198,122876,57130,27364,48020",
        "rating-desc": "Rated 4.3 stars out of five stars\nRated 5 stars out of five stars\nRated 2 stars out of five stars\nRated 3 stars out of five stars\nRated 1 stars out of five stars\nRated 4 stars out of five stars\nRated 4.4 stars out of five stars\nRated 4.1 stars out of five stars\nRated 4.0 stars out of five stars\nRated 4.2 stars out of five stars\nRated 3.7 stars out of five stars\nRated 2.6 stars out of five stars\nRated 4.5 stars out of five stars\nRated 3.9 stars out of five stars",
        "rating-value": "4.276370525360107",
        "software-price": "0",
        "meta-desc": "\r\n Updated November 2, 2015 \r\n Size  15M   \r\n Installs   50,000,000 - 100,000,000   \r\n Current Version  1.3.large   \r\n Requires Android       4.1 and up     \r\n Content Rating Rated for 3+  Learn more  \r\n  Interactive Elements  Shares Location \r\n Permissions  View details  \r\n Report  Flag as inappropriate  \r\n  Offered By  Google Inc. \r\n  Developer    Visit website   Email apps-help@google.com   Privacy Policy  1600 Amphitheatre Parkway, Mountain View 94043  ",
        "file-size": "15",
        "content-rating": "Rated for 3+",
        "date-published": "",
        "software-version": "1.3.large",
        "software-os": "4.1 and up",
        "total-downloads": "75000000",
        "app-url": "https://play.google.com/store/apps/details?id=com.google.android.launcher\u0026hl=en",
        "app-id": "com.google.android.launcher"
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


#### Multiple
```go

$  ./storemeta  -i="544007664,535886823,643496868"
{
        "platform": "IOS",
        "title": "YouTube on the App StoreYouTube\n              By Google, Inc.\n              Essentials\n            \n            \n              \n              View More by This Developer\n              \n            \n            Open iTunes to buy and download apps.",
        "developer": "Google, Inc.",
        "developer-site": "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewEula?id=544007664",
        "genre": "Photo \u0026 Video",
        "description": "Get the official YouTube app for iPhone and iPad.See what the world is watching in music, gaming, entertainment, news and more. Explore the hottest videos on YouTube with the new trending tab. Subscribe to channels, share with friends, edit and upload videos, and watch on any device.With a new design, you can now have fun exploring videos you love more easily and quickly than before. Just tap an icon or swipe to switch between recommended videos, the trending tab, your subscriptions, and your account.You can also subscribe to your favorite channels, create playlists, edit and upload videos, check out what’s trending, express yourself with comments or shares, cast a video to your TV, and more – all from inside the app.With the official YouTube app you can enjoy your favorite videos, creators, and music for free.FIND VIDEOS YOU LOVE FAST- Browse personal recommendations on the home tab- See the latest from the creators you follow on the subscriptions tab- Discover the world’s hottest videos on the trending tab- Look up videos you’ve watched and your like history on the account tabCONNECT AND SHARE- Let people know how you feel with likes, comments, or shares- Edit, add filters, add music, and upload your own videos all inside the app",
        "badge": "This app is designed for both iPhone and iPad",
        "rating-total": "233945",
        "rating-per-star": "723,233222",
        "rating-desc": "3 and a half stars, 723 Ratings",
        "rating-value": "3.39557",
        "software-price": "Free",
        "meta-desc": "Infrequent/Mild Horror/Fear ThemesInfrequent/Mild Profanity or Crude HumorInfrequent/Mild Medical/Treatment InformationInfrequent/Mild Sexual Content and NudityInfrequent/Mild Cartoon or Fantasy ViolenceFrequent/Intense Mature/Suggestive ThemesInfrequent/Mild Realistic ViolenceInfrequent/Mild Alcohol, Tobacco, or Drug Use or ReferencesInfrequent/Mild Simulated Gambling",
        "file-size": "54.2",
        "content-rating": "You must be at least 17 years old to download this app.Infrequent/Mild Horror/Fear ThemesInfrequent/Mild Profanity or Crude HumorInfrequent/Mild Medical/Treatment InformationInfrequent/Mild Sexual Content and NudityInfrequent/Mild Cartoon or Fantasy ViolenceFrequent/Intense Mature/Suggestive ThemesInfrequent/Mild Realistic ViolenceInfrequent/Mild Alcohol, Tobacco, or Drug Use or ReferencesInfrequent/Mild Simulated Gambling",
        "date-published": "2016-08-23 00:00:00",
        "software-version": "11.32",
        "software-os": "Requires iOS 8.0 or later. Compatible with iPhone, iPad, and iPod touch.",
        "total-downloads": "23394500",
        "app-url": "https://itunes.apple.com/us/app/youtube/id544007664?mt=8",
        "app-id": "544007664"
}
{
        "platform": "IOS",
        "title": "Hangouts on the App StoreHangouts\n              By Google, Inc.\n              \n            \n            \n              \n              View More by This Developer\n              \n            \n            Open iTunes to buy and download apps.",
        "developer": "Google, Inc.",
        "developer-site": "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewEula?id=643496868",
        "genre": "SOCIAL-NETWORKING",
        "description": "Use Hangouts to keep in touch. Message friends, start free video or voice calls, and hop on a conversation with one person or a group. Say more with photos, stickers, and emoji.• Include all your friends with group chats for up to 150 people.• Say more with status messages, photos, maps, emoji, stickers, and animated GIFs.• Turn any conversation into a free group video call with up to 10 friends.• Call any phone number in the world (and all calls to other Hangouts users are free!).• Connect your Google Voice account for phone number, SMS, and voicemail integration.• Keep in touch with friends across all your devices.• Message friends anytime, even if they're offline.• Manage mobile app remotely with Google for Work.Note: Mobile carrier and ISP charges may apply. Calls to Hangouts users are free, but other calls might be charged.",
        "badge": "This app is designed for both iPhone and iPad",
        "rating-total": "32476",
        "rating-per-star": "565,31911",
        "rating-desc": "4 and a half stars, 565 Ratings",
        "rating-value": "4.37876",
        "software-price": "Free",
        "meta-desc": "",
        "file-size": "59.5",
        "content-rating": "Rated 4+",
        "date-published": "2016-08-03 00:00:00",
        "software-version": "11.5.0",
        "software-os": "Requires iOS 8.0 or later. Compatible with iPhone, iPad, and iPod touch.",
        "total-downloads": "3247600",
        "app-url": "https://itunes.apple.com/us/app/hangouts/id643496868?mt=8",
        "app-id": "643496868"
}
{
        "platform": "IOS",
        "title": "Chrome - web browser by Google on the App StoreChrome - web browser by Google\n              By Google, Inc.\n              \n            \n            \n              \n              View More by This Developer\n              \n            \n            Open iTunes to buy and download apps.",
        "developer": "Google, Inc.",
        "developer-site": "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewEula?id=535886823",
        "genre": "Utilities",
        "description": "Browse fast on your iPhone and iPad with the Google Chrome browser you love on desktop. Pick up where you left off on your other devices, search by voice, and easily read webpages in any language.• SYNC ACROSS DEVICES - seamlessly access and open tabs and bookmarks from your laptop, phone or tablet• FASTER BROWSING - choose from search results that instantly appear as you type and quickly access previously visited pages• VOICE SEARCH - use the magic of Google voice search to find answers on-the-go without typing• TRANSLATE - easily read webpages in any language• UNLIMITED TABS - open as many tabs as your heart desires and quickly flip through them like a deck of cards• PRIVACY - use Incognito mode to browse without saving your history (learn more at http://goo.gl/WUx02)",
        "badge": "This app is designed for both iPhone and iPad",
        "rating-total": "53485",
        "rating-per-star": "325,53160",
        "rating-desc": "3 and a half stars, 325 Ratings",
        "rating-value": "3.68615",
        "software-price": "Free",
        "meta-desc": "Unrestricted Web Access",
        "file-size": "65.6",
        "content-rating": "You must be at least 17 years old to download this app.Unrestricted Web Access",
        "date-published": "2016-07-27 00:00:00",
        "software-version": "52.0.2743.84",
        "software-os": "Requires iOS 9.0 or later. Compatible with iPhone, iPad, and iPod touch.",
        "total-downloads": "5348500",
        "app-url": "https://itunes.apple.com/us/app/chrome-web-browser-by-google/id535886823?mt=8",
        "app-id": "535886823"
}


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


## Docker Binary

- [x] In order to  use it as dockerize binary


```sh

    sudo  sysctl -w net.ipv4.ip_forward=1

    sudo  docker run --rm  bayugyug/storemeta -h

    sudo  docker run --rm  bayugyug/storemeta -i="293622097"

```

