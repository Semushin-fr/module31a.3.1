;(function (e) {
  function t(t) {
    for (
      var r, u, i = t[0], c = t[1], l = t[2], f = 0, p = [];
      f < i.length;
      f++
    )
      (u = i[f]),
        Object.prototype.hasOwnProperty.call(a, u) && a[u] && p.push(a[u][0]),
        (a[u] = 0)
    for (r in c) Object.prototype.hasOwnProperty.call(c, r) && (e[r] = c[r])
    s && s(t)
    while (p.length) p.shift()()
    return o.push.apply(o, l || []), n()
  }
  function n() {
    for (var e, t = 0; t < o.length; t++) {
      for (var n = o[t], r = !0, i = 1; i < n.length; i++) {
        var c = n[i]
        0 !== a[c] && (r = !1)
      }
      r && (o.splice(t--, 1), (e = u((u.s = n[0]))))
    }
    return e
  }
  var r = {},
    a = { app: 0 },
    o = []
  function u(t) {
    if (r[t]) return r[t].exports
    var n = (r[t] = { i: t, l: !1, exports: {} })
    return e[t].call(n.exports, n, n.exports, u), (n.l = !0), n.exports
  }
  ;(u.m = e),
    (u.c = r),
    (u.d = function (e, t, n) {
      u.o(e, t) || Object.defineProperty(e, t, { enumerable: !0, get: n })
    }),
    (u.r = function (e) {
      'undefined' !== typeof Symbol &&
        Symbol.toStringTag &&
        Object.defineProperty(e, Symbol.toStringTag, { value: 'Module' }),
        Object.defineProperty(e, '__esModule', { value: !0 })
    }),
    (u.t = function (e, t) {
      if ((1 & t && (e = u(e)), 8 & t)) return e
      if (4 & t && 'object' === typeof e && e && e.__esModule) return e
      var n = Object.create(null)
      if (
        (u.r(n),
        Object.defineProperty(n, 'default', { enumerable: !0, value: e }),
        2 & t && 'string' != typeof e)
      )
        for (var r in e)
          u.d(
            n,
            r,
            function (t) {
              return e[t]
            }.bind(null, r)
          )
      return n
    }),
    (u.n = function (e) {
      var t =
        e && e.__esModule
          ? function () {
              return e['default']
            }
          : function () {
              return e
            }
      return u.d(t, 'a', t), t
    }),
    (u.o = function (e, t) {
      return Object.prototype.hasOwnProperty.call(e, t)
    }),
    (u.p = '/')
  var i = (window['webpackJsonp'] = window['webpackJsonp'] || []),
    c = i.push.bind(i)
  ;(i.push = t), (i = i.slice())
  for (var l = 0; l < i.length; l++) t(i[l])
  var s = c
  o.push([0, 'chunk-vendors']), n()
})({
  0: function (e, t, n) {
    e.exports = n('56d7')
  },
  '56d7': function (e, t, n) {
    'use strict'
    n.r(t)
    n('e260'), n('e6cf'), n('cca6'), n('a79d')
    var r = n('2b0e'),
      a = function () {
        var e = this,
          t = e.$createElement,
          n = e._self._c || t
        return n('v-app', [n('v-main', [n('News')], 1)], 1)
      },
      o = [],
      u = function () {
        var e = this,
          t = e.$createElement,
          n = e._self._c || t
        return n(
          'div',
          [
            n('h2', { staticClass: 'mx-5 my-5' }, [
              e._v('GoNews - агрегатор новостей.'),
            ]),
            e._l(e.news, function (t) {
              return n(
                'div',
                { key: t.ID },
                [
                  n(
                    'v-card',
                    {
                      staticClass: 'mx-5 my-5',
                      attrs: { elevation: '10', outlined: '' },
                    },
                    [
                      n('v-card-title', [
                        n('a', { attrs: { href: t.Link, target: '_blank' } }, [
                          e._v(' ' + e._s(t.Title) + ' '),
                        ]),
                      ]),
                      n(
                        'v-card-text',
                        [
                          e._v(' ' + e._s(t.Content) + ' '),
                          n('v-card-subtitle', [
                            e._v(' ' + e._s(new Date(1e3 * t.PubTime)) + ' '),
                          ]),
                        ],
                        1
                      ),
                    ],
                    1
                  ),
                ],
                1
              )
            }),
          ],
          2
        )
      },
      i = [],
      c =
        (n('d3b7'),
        {
          name: 'News',
          data: function () {
            return { news: [] }
          },
          mounted: function () {
            var e = this,
              t = 'http://' + window.location.host + '/news/40'
            fetch(t)
              .then(function (e) {
                return e.json()
              })
              .then(function (t) {
                return (e.news = t)
              })
          },
        }),
      l = c,
      s = n('2877'),
      f = n('6544'),
      p = n.n(f),
      d = n('b0af'),
      v = n('99d9'),
      b = Object(s['a'])(l, u, i, !1, null, '4d65d0c8', null),
      h = b.exports
    p()(b, {
      VCard: d['a'],
      VCardSubtitle: v['a'],
      VCardText: v['b'],
      VCardTitle: v['c'],
    })
    var w = {
        name: 'App',
        components: { News: h },
        data: function () {
          return {}
        },
      },
      y = w,
      m = n('7496'),
      _ = n('f6c4'),
      g = Object(s['a'])(y, a, o, !1, null, null, null),
      O = g.exports
    p()(g, { VApp: m['a'], VMain: _['a'] })
    var j = n('f309')
    r['a'].use(j['a'])
    var x = new j['a']({})
    ;(r['a'].config.productionTip = !1),
      new r['a']({
        vuetify: x,
        render: function (e) {
          return e(O)
        },
      }).$mount('#app')
  },
})
//# sourceMappingURL=app.3e1005d5.js.map
