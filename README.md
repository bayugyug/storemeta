## storemeta


- [x] This is a simple golang script that will parse the ff:

        ** a. Android Play Store meta info
        ** b. Apple App Store meta info 

- [x] Output the meta-info in JSON format



## Compile

```sh

    git clone https://github.com/bayugyug/storemeta.git && cd storemeta && ./storemeta

```



## Usage

```sh


    ./storemeta -a <AndroidStoreID>  -i <IOSStoreID>


    Example:

    ./storemeta  -a="com.google.android.apps.photos"

    or

    ./storemeta  -i="293622097"

    or

    ./storemeta  -a="com.google.android.apps.photos" -i="293622097"


```


## Sample (Android)

```go


    ./storemeta  -a="com.google.android.apps.photos"


    {"platform":"ANDROID","title":"Google Photos - Android Apps on Google Play","developer":"Google Inc.","developer-site":"https://www.google.com/url?q=http://www.google.com/policies/privacy\u0026sa=D\u0026usg=AFQjCNE7y6nm7TcHvct7CDJRmWrYBHvMEQ","genre":"PHOTOGRAPHY","description":"Google Photos is the home for all your photos and videos. Automatically organized and searchable, you can find photos fast and bring them to life. It’s the photo gallery that thinks like you do.  VISUAL SEARCH Your photos are now searchable by the places and things that appear in them. Looking for that fish taco you ate in Hawaii? Just search “food in Hawaii” to find it – no tagging required.  UNLIMITED FREE HIGH QUALITY STORAGE Automatically backup all your photos and videos. Access them on any device or on the web at photos.google.com. Your photos are safe, secure, and private to you.  FREE UP SPACE ON YOUR DEVICE Never worry about running out of space on your phone again. In Settings, just tap “Free up device storage” – photos that are safely backed up will be removed from your device’s storage, but will still be available in Google Photos.  BRING PHOTOS TO LIFE Enjoy automatically created montage movies, interactive stories, collages, animations, panoramas, and more from your photos. Or you can easily create them yourself – just tap +.  EASY EDITING Transform photos with the tap of a finger. Use simple, yet powerful, photo and video editing tools to apply filters, adjust colors, and more.  SHARED ALBUMS Get everyone’s photos and videos in one place, across Android, iOS, and the web. Privately sharing all the photos you took – and getting the ones you didn’t – has never been easier.  INSTANT SHARING Instantly share up to 1,500 photos with anyone – no matter what device they’re on. In the share menu, just tap Create Link.  REDISCOVER YOUR PHOTOS It’s easier than ever to relive your memories. The Assistant can create collages of your old photos that help you relive the past.  READY TO CAST View your photos and videos on your TV with Chromecast support.  Follow us for the latest news and updates Twitter: https://twitter.com/GooglePhotos Google+: https://google.com/+GooglePhotos  Need help? Visit https://support.google.com/photos  Face grouping is not available in all countries.","badge":"Top Developer","rating-total":"4846787","rating-per-star":"3214313,920228,344084,130721,237441","rating-desc":"Rated 4.4 stars out of five stars\nRated 5 stars out of five stars\nRated 4 stars out of five stars\nRated 2 stars out of five stars\nRated 3 stars out of five stars\nRated 1 stars out of five stars\nRated 4.3 stars out of five stars\nRated 4.5 stars out of five stars\nRated 4.2 stars out of five stars\nRated 4.6 stars out of five stars\nRated 3.9 stars out of five stars\nRated 4.1 stars out of five stars\nRated 4.0 stars out of five stars","rating-value":"4.391282558441162","software-price":"0","meta-desc":"\r\n Updated August 22, 2016 \r\n Installs   500,000,000 - 1,000,000,000   \r\n Requires Android       4.0 and up     \r\n Content Rating Rated for 3+  Learn more  \r\n Permissions  View details  \r\n Report  Flag as inappropriate  \r\n  Offered By  Google Inc. \r\n  Developer    Visit website   Email apps-help@google.com   Privacy Policy  1600 Amphitheatre Parkway, Mountain View 94043  ","file-size":"","content-rating":"Rated for 3+","date-published":"","software-version":"","software-os":"4.0 and up","total-downloads":"750000000","app-url":"https://play.google.com/store/apps/details?id=com.google.android.apps.photos\u0026hl=en","app-id":"com.google.android.apps.photos"}

```

## Sample (IOS)

```go


    ./storemeta  -i="293622097"


    {"platform":"IOS","title":"Google Earth on the App StoreGoogle Earth\n              By Google, Inc.\n              \n            \n            \n              \n              View More by This Developer\n              \n            \n            Open iTunes to buy and download apps.","developer":"Google, Inc.","developer-site":"https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewEula?id=293622097","genre":"TRAVEL","description":"Fly around the planet with a swipe of your finger with Google Earth for iPhone, iPad, and iPod touch. Explore distant lands or reacquaint yourself with your childhood home. Search for cities, places, and businesses. Browse layers including roads, borders, places, photos and more.  See the world at street level with integrated Street View.Use the “tour guide” to easily discover exciting new places to explore. With a quick swipe on the tab at the bottom of the screen, you can bring up a selection of virtual tours from around the globe.With 3D imagery, you can now fly through complete 3D recreations of select cities, including San Francisco, Boston, Rome, and others. With every building modeled in 3D, you truly get a sense of flying above the city.","badge":"This app is designed for both iPhone and iPad","rating-total":"446047","rating-per-star":"378,445669","rating-desc":"3 and a half stars, 378 Ratings","rating-value":"3.69312","software-price":"Free","meta-desc":"","file-size":"30.1","content-rating":"Rated 4+","date-published":"2016-05-03 00:00:00","software-version":"7.1.6","software-os":"Requires iOS 5.0 or later. Compatible with iPhone, iPad, and iPod touch.","total-downloads":"44604700","app-url":"https://itunes.apple.com/us/app/google-earth/id293622097?mt=8","app-id":"293622097"}

```


