# GoHoneywellAPI

basic implementation to use the Honeywell API for *T* compatible devices (T9, T10, Lyric)

## Usage

First you need to register on the Honeywell API portal at <https://developer.honeywellhome.com/user/register>

Once done, you have to create a new application. This will give you access to a `Consumer Key` and a `Consumet Secret`
![Honeywell app creation](https://raw.githubusercontent.com/prune998/gohoneywellapi/master/docs/pictures/hwApp.png)

You also need to register your devices (thermostat) on the Honeywell Application, iPhone or Android. You will be asked to create an account, which is different from the one you juste created for the API.

Following the [oAuth2 guide](https://developer.honeywellhome.com/content/oauth2-guide) you will then need to open an URL in your browser.
You will have to authenticate using the login from the mobile App, not the one from the developper API !

The URL is like:

```https://api.honeywell.com/oauth2/authorize?response_type=code&client_id={apikey}&redirect_uri={redirectUri}```

The `apikey` here is your `Consumer Key`.
Use `none` as the `redirect URL`.

If your `Consumer Key` is like `lknwcoiwenerkfr`, then your final URL is :

```https://api.honeywell.com/oauth2/authorize?response_type=code&client_id=lknwcoiwenerkfr&redirect_uri=none```

As I said, authenticate, and allow your application to access your devices by answering `Allow`.
![Honeywell App Authz](https://raw.githubusercontent.com/prune998/gohoneywellapi/master/docs/pictures/goT9_authz.png)

The webpage will redirect you to some URL like 

```https://api.honeywell.com/oauth2/app/none?code=qwerty&scope=```

Note the code down (`qwerty` in this example). You will need it along your Key and Secret to access the API and get your Bearer Token.

## References

1. Honeywell Developper site : <https://developer.honeywellhome.com/content/getting-started-guide>
1. Using the API by hand : <https://www.domoticz.com/wiki/HoneywellDeveloperApiKey>
