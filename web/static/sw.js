importScripts('/_nuxt/workbox.4c4f5ca6.js')

workbox.precaching.precacheAndRoute([
  {
    "url": "/_nuxt/052babc9643430a799ae.js",
    "revision": "503857aa4a0b220131cbbd8c49d2df66"
  },
  {
    "url": "/_nuxt/2600e5825b40c6385983.js",
    "revision": "d028668e341040ab85bc64e5077e5907"
  },
  {
    "url": "/_nuxt/4c87ffc4dace39b28b3f.js",
    "revision": "b38ef858b5798dfd4af72079a18fafb3"
  },
  {
    "url": "/_nuxt/62ae0ce440767aaae95b.js",
    "revision": "d3f7dbac1d430b37d07f5fe0b8feea18"
  },
  {
    "url": "/_nuxt/e66c3a709a9ed71061cb.js",
    "revision": "be70ac9a380f95fb9a4accfe9f4607df"
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
