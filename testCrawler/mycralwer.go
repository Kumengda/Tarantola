package testCrawler

import (
	"fmt"
	"github.com/Kumengda/Tarantola/tarantola"
)

type MyCrawler struct {
	tarantola.BaseCrawler
}

func (m *MyCrawler) Crawl() error {
	jsCode := `
T = 'replace';
function test(result) {
  result = JSON.parse(result)
  var n;
  f= false;
  Zt = 'params';
  N2 = 'hasOwnProperty';
  i7 = 'keys';
  Z = 'Object';
  M = 'forEach';
  a = [];
  p = 'analysis';
  Ot = 'sort';
  I1 = 'join';
  _ = '';
  v = '@#';
  Jt = 'url';
  Mt = 'baseURL';
  r = 45545981687;
  d = 'xyz517cda96efgh';
  j = 'index';
  B = 1;
  Rn = "?";
  Hn = "&";
  B1 = "=";
  V1 = "encodeURIComponent";
  H = 0;
  var F;
  f || F != s
  s = 21537;
  n = '1706764581.077';
  b = 'push';
  var e, r = +new Date() - (s || H) - 1661224081041, a = [];
  a.push(result[Zt]['word'])
  a = a.sort().join("")
  a = v2(a)

  a = (a += v + result[Jt][T](result[Mt], _)) + (v + r) + (v + 3)
  e = v2(h2(a,d))
  return encodeURIComponent(e);
}

function s(n) {
  var n = new z[M1](T1 + n + C1);
  return (n = z[I][P1][A1](n)) ? z[g2](n[2]) : F
}
function m(n, t, e) {
  var r = new z[W];
  r[z1](r[R1]() + e),
    z[I][P1] = n + B1 + z[H1](t) + (F == e ? _ : _1 + r[E1]()) + N1
}
function l(n) {
  return z[G1](function(t) {
    try {
      return z[F1](t)
    } catch (n) {
      return z[W1][K1](t, Z1)[U1]()
    }
  }(n)[$1](_)[D1](function(n) {
    return O1 + (J1 + n[q1](H)[U1](16))[j1](-2)
  })[I1](_))
}
function v2(t) {
  Y1 = '0x';
  t = encodeURIComponent(t).replace(/%([0-9A-F]{2})/g, function(n, t) {
    return o(Y1 + t)
  });
  return btoa(t)
}

function o(n) {

  t = "",
    ['66', '72', "6f", "6d", "43", "68", "61", "72", "43", "6f", "64", "65"].forEach(function(n) {
      t += unescape("%u00" + n)
    });
  var t, e = t;
  return String.fromCharCode(n)
}
function p(n, t) {
  t = t || u();
  for (var e = (n = n[$1](_))[R], r = t[R], a = q1, i = H; i < e; i++)
    n[i] = o(n[i][a](H) ^ t[i % r][a](H));
  return n[I1](_)
}
function h2(n, t) {
  R = "length";
  q1= "charCodeAt";
  for (var e = (n = n['split'](_))[R], r = t[R], a = q1, i = H; i < e; i++)
    n[i] = o(n[i][a](H) ^ t[(i + 10) % r][a](H));
  return n[I1](_)
}
function y(n, t, e) {
  for (var r = void 0 === e ? 2166136261 : e, a = H, i = n[R]; a < i; a++)
    r = (r ^= n[q1](a)) + ((r << B) + (r << 4) + (r << 7) + (r << 8) + (r << 24));
  return t ? (n3 + (r >>> H)[U1](16) + t3)[X1](-16) : r >>> H
}
function g(n, t) {
  var e, r, a, i, o = 2 < arguments[R] && void 0 !== arguments[2] ? arguments[2] : 3, u = (e3 != typeof z[r3] && (e = z[S],
    i = z[I],
    u = a3,
    e[r = r3] = e[r] || function() {
      (e[r][i3] = e[r][i3] || [])[b](arguments)
    }
    ,
    a = i[G](u),
    i = i[o3](u)[H],
    a[u3] = !H,
    a[c3] = f3,
    a[s3] = d3,
    i[l3][m3](a, i)),
    v3);
  z[r3](p3, u),
    z[r3](h3),
    z[r3](y3, o),
  S2 != (0,
    c[E])(n) && (n = {}),
    n = z[Z][g3]({
      "\u7528\u6237\u540d": w3,
      "\u90ae\u7bb1": _,
      "\u6295\u653e": _,
      "\u5173\u6ce8": _
    }, n),
    z[r3](b3, n),
    z[r3](k3, t)
}
test('{"url":"/search/autoCompleteCompany","method":"get","headers":{"common":{"Accept":"application/json, text/plain, */*"},"delete":{},"get":{},"head":{},"post":{"Content-Type":"application/x-www-form-urlencoded"},"put":{"Content-Type":"application/x-www-form-urlencoded"},"patch":{"Content-Type":"application/x-www-form-urlencoded"}},"params":{"word":"哈哈"},"baseURL":"https://api.qimai.cn","transformRequest":[null],"transformResponse":[null],"timeout":15000,"withCredentials":true,"xsrfCookieName":"XSRF-TOKEN","xsrfHeaderName":"X-XSRF-TOKEN","maxContentLength":-1,"maxBodyLength":-1}')
`
	chrome, err := m.ExecJsWithChrome(jsCode)
	if err != nil {
		return err
	}
	fmt.Println(chrome)
	panic("implement me")
}

func NewMyCrawler(options tarantola.BaseOptions) *MyCrawler {
	return &MyCrawler{
		BaseCrawler: tarantola.BaseCrawler{
			BaseOptions: options,
		},
	}
}
