importScripts('/_nuxt/workbox.4c4f5ca6.js')

workbox.precaching.precacheAndRoute([
  {
    "url": "/_nuxt/0e4c68ee0f6909f62d59.js",
    "revision": "11a71e6c6ae1bd1c836930b768cd250a"
  },
  {
    "url": "/_nuxt/1b99f7a7d7df3c8225e7.js",
    "revision": "03dce42180b0884eecded9871f332ed4"
  },
  {
    "url": "/_nuxt/854208b323274f737129.js",
    "revision": "280ecc92c9691a2891084fa53d962fbe"
  },
  {
    "url": "/_nuxt/a9bdbc18f8ed881f2341.js",
    "revision": "7f1d250523c5437a26530e82734f7e45"
  },
  {
    "url": "/_nuxt/c620bf67f9b978630f07.js",
    "revision": "9c2991a07574d4d390aa5c019b1ead00"
  }
], {
  "cacheId": "web",
  "directoryIndex": "/",
  "cleanUrls": false
})

workbox.clientsClaim()
workbox.skipWaiting()

workbox.routing.registerRoute(new RegExp('/_nuxt/.*'), workbox.strategies.cacheFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('/.*'), workbox.strategies.networkFirst({}), 'GET')
