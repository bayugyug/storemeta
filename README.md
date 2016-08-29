## storemeta


- [x] This is a simple golang script that will parse the ff:
            a. Android Play Store meta info
            b. Apple App Store meta info 

- [x] Output the meta-info in JSON format


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
