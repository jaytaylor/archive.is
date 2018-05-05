package archiveis

import (
	"strings"
	"testing"
)

func TestCheckCrawlResult(t *testing.T) {
	testCases := map[string]string{
		`<html><body>
                           <div>
      <img width="48" height="48" style="vertical-align:middle" src="https://archive.is/alert_error.gif"/>
      <span style="vertical-align:middle;font-size:48px;padding-left:5px">Error: Network error.</span>
      <hr/>
    </div>
                           <table cellspacing="0" cellpadding="0" border="0" style="font-family: monospace; font-size: 10px">
              <tr align="left">
                <th>status</th>
                <th style="text-align:left;padding-left:1em">type</th>
                <th style="text-align:right;padding-left:1em">size</th>
                <th></th>
                <th style="text-align:left;padding-left:1em">url</th>
              </tr>
              <tr valign="top" style="background-color:#FFFFFF">
                    <td style="text-align:right;padding-left:1em" colspan="3">
                      
                    </td>
                    <td style="padding-left:1em">GET</td>
                    <td style="padding-left:1em">
                      <a target="_blank" style="text-decoration:none;word-wrap:break-word;word-break:break-all" href="http://uscode.house.gov/download/releasepoints/us/pl/115/168not141/pdf_usc01@115-168not141.zip">http://uscode.house.gov/download/releasepoints/us/pl/115/168not141/pdf_usc01@115-168not141.zip</a>
                    </td>
                  </tr>
            </table>
                         </body></html>`: "Network Error",

		`<!DOCTYPE html><html style="background-color:#EEEEEE" prefix="og: http://ogp.me/ns# article: http://ogp.me/ns/article#" itemscope itemtype="http://schema.org/Article"><!--73.231.253.6--><!--Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36--><head><meta http-equiv="Content-Type" content="text/html;charset=utf-8"/><meta name="robots" content="index,noarchive"/><meta property="twitter:card" content="summary"/><meta property="twitter:site" content="@archiveis"/><meta property="og:type" content="article"/><meta property="og:site_name" content="archive.is"/><meta property="og:url" content="http://archive.is/XvEff" itemprop="url"/><meta property="og:title" content="TXT-Web"/><meta property="twitter:title" content="TXT-Web"/><meta property="twitter:description" content="archived 5 May 2018 22:28:55 UTC" itemprop="description"/><meta property="article:published_time" content="2018-05-05T22:28:55Z" itemprop="dateCreated"/><meta property="article:modified_time" content="2018-05-05T22:28:55Z" itemprop="dateModified"/><link rel="image_src" href="https://archive.is/XvEff/0fcfd78d0f8b8379e1b8aa5e4fb7743aa22268d4/scr.png"/><meta property="og:image" content="https://archive.is/XvEff/0fcfd78d0f8b8379e1b8aa5e4fb7743aa22268d4/scr.png" itemprop="image"/><meta property="twitter:image" content="https://archive.is/XvEff/0fcfd78d0f8b8379e1b8aa5e4fb7743aa22268d4/scr.png"/><meta property="twitter:image:src" content="https://archive.is/XvEff/0fcfd78d0f8b8379e1b8aa5e4fb7743aa22268d4/scr.png"/><meta property="twitter:image:width" content="1024"/><meta property="twitter:image:height" content="768"/><link rel="icon" href="//www.google.com/s2/favicons?domain=txt.gigawatt.io"/><link rel="canonical" href="https://archive.is/XvEff"/><link rel="bookmark" href="http://archive.today/20180505222855/https://txt.gigawatt.io/"/><title>TXT-Web</title></head><body style="margin:0;background-color:#EEEEEE"><center><div id="HEADER" style="font-family:sans-serif;background-color:#FFFAE1;border-bottom:2px #B40010 solid;min-width:1028px"><div style="padding-top:10px"></div><table style="width:1028px;font-size:10px" border="0" cellspacing="0" cellpadding="0"><tr><td style="width:150px;text-align:center;vertical-align:top" rowspan="2"><a style="text-decoration:none;white-space:nowrap;color:black;margin:0px;cursor:pointer" href="https://archive.today/"><div style="font-size:24px">archive.today</div><div style="font-size:12px">webpage capture</div></a></td><td style="text-align:right;padding:3px 3px 0 3px;white-space:nowrap;vertical-align:top;font-size:14px;font-weight:bold">Saved from</td><td style="text-align:right;padding:3px 3px 0 3px;white-space:nowrap;vertical-align:top"><form style="text-align:left;margin:0" action="https://archive.is/search/" method="get"><table cellspacing="0" cellpadding="0" border="0"><tr><td style="width:500px"><input style="border:1px solid black;height:20px;margin:0 0 0 0;padding:0;width:500px" type="text" name="q" value="https://txt.gigawatt.io/"/><input type="hidden" name="t" value="1525559335681"/><input type="hidden" name="id" value="XvEff"/><div style="text-align:right;font-size:10px"><a style="display:inline-block;white-space:nowrap;padding:0;margin:0 4px;text-decoration:underline;color:#1D2D40" href="https://archive.is/https://txt.gigawatt.io/" title=""><img style="width:16px;height:16px;border:0;margin-top:-3px;position:relative;padding-right:2px;top:5px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAMAAAAoLQ9TAAAAflBMVEVMaXFLS0tLS0tLS0tLS0tLS0tMTExLS0tLS0tMTExMTExLS0tLS0tLS0tLS0tLS0tLS0tMTExMTExLS0tMTExLS0tPT09LS0tLS0tLS0tLS0tLS0tLS0tLS0tPT09PT09LS0tPT09LS0tPT09PT09LS0tLS0tLS0tLS0tPT0+yjOdVAAAAKHRSTlMAi0F+99cB+/0SK9oZvgZh7WWH4AkcPuiauMoL0MnvNU3PWk5qsoAkVcfFqAAAAJBJREFUeF5tj+kOgzAMg12gB1AYlHPnfaTv/4JbUyFtEv4RpZ/k2MW6irkzWjqLKOUMZVJ+xyYCQU3bA3mqiUlC0xvwYdPZAKCiCyJASgLYm04tIDcltjtqqgJnXwRSaggiOtYY/antGaiKpgFQ96t/sAX5IQH73chHg+oQFmMtv5//xaBeS/Xb7+cyKSxW9AH2lQlEdnL3oAAAAABJRU5ErkJggg=="/>history</a></div></td><td style="vertical-align:top"><input style="width:60px;height:20px;padding:0;margin:0 0 0 3px" type="submit" tabindex="-1" value="search"/></td></tr></table></form></td><td style="text-align:right;padding:4px 5px 2px 5px;font-size:14px;white-space:nowrap;vertical-align:top" rowspan="1"><time itemprop="pubdate" datetime="2018-05-05T22:28:55Z">5 May 2018 22:28:55 UTC</time></td></tr><tr><td style="text-align:right;font-size:12px;padding:0 3px 3px 3px;vertical-align:top;font-weight:bold;white-space:nowrap" colspan="1">All snapshots</td><td style="text-align:left;font-size:12px;padding:0 3px 3px 3px;vertical-align:top" colspan="2" rowspan="2"><b>from host </b><a style="color:#1D2D40" href="https://archive.is/txt.gigawatt.io">txt.gigawatt.io</a></td></tr><tr><td style="vertical-align:bottom;text-align:left;white-space:nowrap" colspan="2" rowspan="2"><a style="margin-left:20px;position:relative;top:2px;text-decoration:none;display:inline-block;vertical-align:middle;padding:0 10px;line-height:24px;;font-size:12px;font-weight:bold;background-color:#EEEEEE;border-width:2px 2px 0px 2px;border-style:solid;border-color:#B40010 #B40010 #FFFAE1 #B40010;color:black" href="https://archive.is/XvEff">Webpage</a><a style="margin-left:10px;position:relative;top:2px;text-decoration:none;display:inline-block;vertical-align:middle;padding:0 10px;line-height:24px;;font-size:12px;background-color:#B40010;border-width:2px 2px 0px 2px;border-style:solid;border-color:#B40010;color:white" href="https://archive.is/XvEff/image">Screenshot</a></td></tr><tr><td style="text-align:right;padding:0 5px 3px 0;white-space:nowrap"><a style="line-height:16px;vertical-align:middle;margin-right:20px;text-decoration:underline;color:#1D2D40" href="https://archive.is/XvEff/share" onclick="return showDivShare()"><img style="width:16px;height:16px;border:0;position:relative;top:5px;padding-right:2px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAMAAAAoLQ9TAAABU1BMVEUAAAABAQEAAAAAAAAAAAAAAAAAAAAAAAAAAAALCwsfHx8dHR0ICAgCAgIVFRUSEhISEhIODg5TU1MFBQUCAgIDAwNJSUkTExMEBAQHBwcICAgEBAQMDAwCAgJZWVk5OTkKCgoUFBQhISEODg4KCgoAAAANDQ0MDAwCAgIBAQEjIyMCAgIvLy8AAAAVFRUPDw8ODg4QEBAAAAAKCgovLy8qKioKCgoTExMAAAASEhIODg4EBAQAAAAFBQUKCgoKCgoAAAAAAAAAAAAAAABcXFwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQEBAAAAAeHh4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAgIoKCgkJCQNDQ12dnYSEhIdHR0lJSUmJiYwMDAUFBQZGRkKCgpiYmKCFFTcAAAAYnRSTlMAAAQCB+9A/gFZYMtZCwADOK7fSgEA9k8APVM3pCfZpCenzKzqGgAAABHgJpHABxMZeE4k4epZw4I8eK5zAF1oSfMhmf7HKzGF1tO0RzIxna+4zTUDJecM4KyL0GEQw2Xe1BogW7kAAAC/SURBVHheNY1Tg8RAEAY7yWZt2zbOtm2jszji/3OnZy710l/VSwOjGQbCYrPTkSRrMuWQ2Mp85ig4Xd7Rn3uRLZ+nRiHgr28EQ3IkGosnQGA+Sp9kB+/f+YLwYqlcqfZ/fjVsGHhotaHz8aYho0vem5icmh4iZ4bC7Fd/bh4FCxSUpWVYEb66BvqfdXLT5ta2rKed3b39g0PVdKxwNZ6enV8AXCJe8c/XN4i37Mp39w9APKqIT3wpIHh+eTXCP2MHbSZMouN0pwAAAABJRU5ErkJggg=="/>share</a><a style="line-height:16px;vertical-align:middle;margin-right:20px;text-decoration:underline;color:#1D2D40" href="https://archive.is/download/XvEff.zip">download .zip</a><a style="line-height:16px;vertical-align:middle;text-decoration:underline;color:#1D2D40" href="https://archive.is/XvEff/abuse">report error or abuse</a></td><td style="text-align:right;padding:0 5px 3px 0;white-space:nowrap"><div id="DIVSHARE" style="position:fixed;padding:70px 50px 50px 50px;top:0;left:0;right:0;bottom:0;background-color:rgba(0,0,0,0.3);z-index:1000000001;display:none" onclick="if (event.target.id=='DIVSHARE' || event.target.id=='DIVSHARE2') { this.style.display='none' }"><center id="DIVSHARE2"><div style="display:table;padding:20px;background-color:#FFFAE1;border:#B40010 5px solid;text-align:left;position:relative"><div style="position:absolute;top:5px;right:5px"><a href="https://archive.is/XvEff" onclick="document.getElementById('DIVSHARE').style.display='none'; return false"><img id="SHARE_CLOSEBTN" style="width:24px;height:24px;opacity:0.3;cursor:pointer" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAMAAADXqc3KAAAATlBMVEUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADEoqZZAAAAGXRSTlMAAhXQyPf4D9gbEvYMwwfu582+tXa7gG0c7UlRPQAAAKtJREFUKM99klcSgzAMRLExJnEhoaTs/S8alB1cCEFf2vfwjKShOSvVKFXFlIL0GYete/lJi2Fyc7cwhTvQroa8N8DN8asHPA25R2xYukX3NcLX9nIlV8kkrmrjQuaFsTARNvHK7Hga01oOV4shihj7owcAn/zw0ezMtpd2MnU2xb66MnlflQxrhuf83MdjEsqzk9PI2dkvMO/ibhHPLbmhHHFwhz+D9Ex/6gNJowqlzHFZ/gAAAABJRU5ErkJggg=="/></a></div><div style="display:table-row"><div style="display:table-cell;padding:5px;vertical-align:top"></div><div style="display:table-cell;padding:5px;width:600px"><div style="display:table-row"><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px  -907px" title="Share to Reddit" onclick="window.open('http://reddit.com/submit?' + 'title=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">Reddit</button></div><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAABDlBMVEX////j4+Pi4uL8/f0nWYT7+/yHh4fg4OAJCQnZ2dnFxcX29vbPz8/e3t6Ojo76+vr3+Pjr6+vm5ubW1tbBwcGYmJj5+fnx8fHu7u/e5ezc3eDb29vKysq3t7ezs7N4lq+ioqKRkZEuX4gqW4UxJXY/Pz/z8/Ps7Ozo6Ojl5eXU3OLR0dGtvs3Hx8ebsMOMqL5vka2mpqZUfJ5FcJQ0KXceHh7t8fXo7fLn7fHZ4enW4OjK1uHD0d3T2NvN09jR0dXEwdC6t8q3tMi7u7uUqbuCn7iTjrKwsLCoqKiEfqhjh6afn59cgZ9Zf55NdZeTk5NWTIyEhIR/f399fX0qHnFtbW1UVFRQUFAvLy8QEBBa7Md2AAABhklEQVQ4y3WSh26DMBRFnwfGzOykzWxWd7OadO+99/z/H6lNcYCEHCTQ0zu6F5DBZx+irM5H573GLhD/kmiDXlTo0TsSntdpMxIx36T0AIKE5CaltbBwSSkdaEA0iVA6Ym7sBPvtBm0+1VbTyGdlvSYiuoHQfegcogmOLh631J4UUSxzSrBQPGkVYMwQMAKPjQqawYIu/1mWVeVgf306alFpvyJJzi0B9Bm7kcM353Ul/HB+LR66y1gWGGM5JGhz/qa6R5wviGeGCaDv5rJe3+/IVAnvvO1IsVqvZwE0AikkWMugMbYu70kgGngkURzLEJCKfL6PBQH5GCENYeamhaWIQNKTguFYobV8z6iglzFKWQkyVlSJXizqUrBt7IFUApBluT8dtlq3ssA08D9BScKLf/kYHouCjTIOCyQ4NyetK1lgruGJBCktSuPMKzCd6QRlIJwqmDaOSRAkDIyc8wQYpahAxk7VzbnPBdBWZiSAWcmU9QJAvjSRoO5Leh4seQjuM77wBxglHL84LRjiAAAAAElFTkSuQmCC) 2px 2px no-repeat" title="Share to Voat" onclick="window.open('http://voat.co/submit?' + 'linkpost=true&amp;title=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">Voat</button></div><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px -1003px" title="Share to Twitter" onclick="window.open('https://twitter.com/intent/tweet?'                                          + '&amp;status=' + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value) + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">Twitter</button></div></div><div style="display:table-row"><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px   -70px" title="Share to VKontakte" onclick="window.open('http://vk.com/share.php?' + 'title=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">VKontakte</button></div><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAgVBMVEUAWun///8KYeoCXOn9/v7N3/sode0VaOu00Pr5+/7n8P72+f7v9f7k7f3Q4fxqnvJAhfA3f+8zfe+ZvfeQuPaArPR1pvIud+0RZeoHXun6+/7c6fzX5fzD2Pq91fqiw/ibvvaFsfZwo/R3pvNZlfJclfFRjvBJh+8Yauwgbusaauuxv4pvAAAAo0lEQVQ4y9WRRxLDIAwABbIJuOJup/f2/weGDBwRXJO97g6DJPgxeMSvJA/7tMKwZ4eIZ3mNIW9Iegx6U8jZP5+sZGKLSvgLxN4WLNMo6gE9zZ5Z8p0qGvDw3DBH1oKXqXC+AYKr+nqlgeRkPrpmHX2SZWve75KRDIQqW8Ahv1HBuRRgqIsHEYxujfeSGBTBoVMBYSb5jhSvI48U8yVW8AX+hg8efQYxRxp+awAAAABJRU5ErkJggg==) 2px 2px no-repeat" title="Share to Pinboard" onclick="window.open('https://pinboard.in/add?' + 'next=same&amp;title=TXT-Web&amp;description=archived+5+May+2018+22:28:55+UTC' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">Pinboard</button></div><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px  -691px" title="Share to Livejournal" onclick="window.open('http://www.livejournal.com/update.bml?' + 'subject=TXT-Web' + '&amp;event='  + encodeURIComponent(document.getElementById('SHARE_HTMLCODE' ).value), '_blank', 'height=650,width=1024,scrollbars=1')">Livejournal</button></div></div><div style="display:table-row"><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px  -286px" title="Share to Facebook" onclick="window.open('http://www.facebook.com/sharer/sharer.php?' +  + '&amp;u='      + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">Facebook</button></div><div style="display:table-cell;padding-right:15px;padding-bottom:10px"><button style="width:180px;height:40px;padding-left:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px -1311px" title="Share to Google+" onclick="window.open('https://plus.google.com/share?' + 't=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')">Google+</button></div></div></div></div><div style="display:table-row"><div style="display:table-cell;padding:5px;vertical-align:top">short link</div><div style="display:table-cell;padding:5px"><input id="SHARE_SHORTLINK" style="width:600px" value="http://archive.today/XvEff"/></div></div><div style="display:table-row"><div style="display:table-cell;padding:5px;vertical-align:top">long link</div><div style="display:table-cell;padding:5px"><input id="SHARE_LONGLINK" style="width:600px" value="http://archive.today/2018.05.05-222855/https://txt.gigawatt.io/"/></div></div><div style="display:table-row"><div style="display:table-cell;padding:5px;vertical-align:top">markdown</div><div style="display:table-cell;padding:5px"><input id="SHARE_MARKDOWN" style="width:600px" value="[archive.today link](http://archive.today/XvEff)"/></div></div><div style="display:table-row"><div style="display:table-cell;padding:5px;vertical-align:top">html code</div><div style="display:table-cell;padding:5px"><textarea id="SHARE_HTMLCODE" style="width:600px;height:100px" wrap="off">&lt;a href=&quot;http://archive.today/XvEff&quot;&gt;
 &lt;img style=&quot;width:300px;height:200px;background-color:white&quot; src=&quot;https://archive.is/XvEff/0fcfd78d0f8b8379e1b8aa5e4fb7743aa22268d4/scr.png&quot;&gt;&lt;br&gt;
 TXT-Web&lt;br&gt;
 archived 5 May 2018 22:28:55 UTC
&lt;/a&gt;</textarea></div></div><div style="display:table-row"><div style="display:table-cell;padding:5px;vertical-align:top">wiki code</div><div style="display:table-cell;padding:5px"><textarea id="SHARE_WIKICODE" style="width:600px;height:100px" wrap="off">{{cite web
 | title       = TXT-Web
 | url         = https://txt.gigawatt.io/
 | date        = 2018-05-05
 | archiveurl  = http://archive.today/XvEff
 | archivedate = 2018-05-05 }}</textarea></div></div></div></center></div></td></tr></table></div><div style="padding:10px 0;min-width:1028px;background-color:#EEEEEE"></div><div id="SOLID" style="background-color:#EEEEEE;padding-bottom:15px"><div id="SHARER" style="position:absolute;width:360px;height:80px;display:none;z-index:1000000000;background-color:#FFFAE1;border:#B40010 5px solid;box-shadow:10px 20px 30px #1D2D40"><center><div style="display:inline-block;padding-left:8px;padding-top:16px;padding-bottom:16px;padding-right:8px"><button style="width:40px;height:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px  -907px" title="Share to Reddit" onclick="window.open('http://reddit.com/submit?' + 'title=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')"></button></div><div style="display:inline-block;padding-left:8px;padding-top:16px;padding-bottom:16px;padding-right:8px"><button style="width:40px;height:40px;text-align:left;font-weight:bold;background:#ccc url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAABDlBMVEX////j4+Pi4uL8/f0nWYT7+/yHh4fg4OAJCQnZ2dnFxcX29vbPz8/e3t6Ojo76+vr3+Pjr6+vm5ubW1tbBwcGYmJj5+fnx8fHu7u/e5ezc3eDb29vKysq3t7ezs7N4lq+ioqKRkZEuX4gqW4UxJXY/Pz/z8/Ps7Ozo6Ojl5eXU3OLR0dGtvs3Hx8ebsMOMqL5vka2mpqZUfJ5FcJQ0KXceHh7t8fXo7fLn7fHZ4enW4OjK1uHD0d3T2NvN09jR0dXEwdC6t8q3tMi7u7uUqbuCn7iTjrKwsLCoqKiEfqhjh6afn59cgZ9Zf55NdZeTk5NWTIyEhIR/f399fX0qHnFtbW1UVFRQUFAvLy8QEBBa7Md2AAABhklEQVQ4y3WSh26DMBRFnwfGzOykzWxWd7OadO+99/z/H6lNcYCEHCTQ0zu6F5DBZx+irM5H573GLhD/kmiDXlTo0TsSntdpMxIx36T0AIKE5CaltbBwSSkdaEA0iVA6Ym7sBPvtBm0+1VbTyGdlvSYiuoHQfegcogmOLh631J4UUSxzSrBQPGkVYMwQMAKPjQqawYIu/1mWVeVgf306alFpvyJJzi0B9Bm7kcM353Ul/HB+LR66y1gWGGM5JGhz/qa6R5wviGeGCaDv5rJe3+/IVAnvvO1IsVqvZwE0AikkWMugMbYu70kgGngkURzLEJCKfL6PBQH5GCENYeamhaWIQNKTguFYobV8z6iglzFKWQkyVlSJXizqUrBt7IFUApBluT8dtlq3ssA08D9BScKLf/kYHouCjTIOCyQ4NyetK1lgruGJBCktSuPMKzCd6QRlIJwqmDaOSRAkDIyc8wQYpahAxk7VzbnPBdBWZiSAWcmU9QJAvjSRoO5Leh4seQjuM77wBxglHL84LRjiAAAAAElFTkSuQmCC) 2px 2px no-repeat" title="Share to Voat" onclick="window.open('http://voat.co/submit?' + 'linkpost=true&amp;title=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')"></button></div><div style="display:inline-block;padding-left:8px;padding-top:16px;padding-bottom:16px;padding-right:8px"><button style="width:40px;height:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px -1003px" title="Share to Twitter" onclick="window.open('https://twitter.com/intent/tweet?'                                          + '&amp;status=' + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value) + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')"></button></div><div style="display:inline-block;padding-left:8px;padding-top:16px;padding-bottom:16px;padding-right:8px"><button style="width:40px;height:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px   -70px" title="Share to VKontakte" onclick="window.open('http://vk.com/share.php?' + 'title=TXT-Web' + '&amp;url='    + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')"></button></div><div style="display:inline-block;padding-left:8px;padding-top:16px;padding-bottom:16px;padding-right:8px"><button style="width:40px;height:40px;text-align:left;font-weight:bold;background:#ccc url(https://archive.is/DpB37/7d07f9e412d40fbb6a6956cffc931e5e5c59e8fd.png) no-repeat 2px  -286px" title="Share to Facebook" onclick="window.open('http://www.facebook.com/sharer/sharer.php?' +  + '&amp;u='      + encodeURIComponent(document.getElementById('SHARE_SHORTLINK').value), '_blank', 'height=650,width=1024,scrollbars=1')"></button></div><div style="display:inline-block;padding-left:8px;padding-top:16px;padding-bottom:16px;padding-right:8px"><button style="width:40px;height:40px;text-align:left;font-weight:bold;background:#ccc url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAPFBMVEUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADQLyYwAAAAE3RSTlMA5gb536akmIyJHBkOAb8zMsHAii9UKQAAAFhJREFUOMvtjjkOwDAIBAHfVy7//68BCneJa0uehhFaaRc2i0HJoi8sNSCGylIc2kQjEDtzEDQn4hrQLRJHALtwQTZyTYZTHzgNfFR4rni04m+k15GbtXgBC1EF1n9jtm4AAAAASUVORK5CYII=) 2px 2px no-repeat" title="More" onclick="showDivShare()"></button></div></center></div><div id="CONTENT" style="background-color:white;min-height:768px;max-height:1000000px;position:relative;border:2px #999999 solid;margin:0px -2px;width:1024px" onclick=""><div class="html1" style="width: 1024px;text-align: left;overflow-x: auto;overflow-y: auto; background-color: rgba(0, 0, 0, 0);position: relative;min-height: 768px;; z-index: 0"><div class="html" style="text-align:left;overflow-x:visible;overflow-y:visible;">



<meta name="description" content="Turn any webpage into plaintext"/>

<div class="body" style="vertical-align:bottom;min-height:115px;color:rgb(0, 0, 0);text-align:left;overflow-x:visible;overflow-y:visible;margin: 8px; ">
Welcome to TXT-Web! <br style="text-align:left;"/>
<br style="text-align:left;"/>
Turn any page into plaintext. <br style="text-align:left;"/>
<br style="text-align:left;"/>
<form style="text-align:left;" action="https://archive.today/XvEff#" method="post">
<input value="" style="text-align:left;"/>
<button>Go</button>
</form>


</div></div></div><!--[if !IE]><!--><div style="position:absolute;right:1028px;top:-14px;bottom:-2px"><table id="hashtags" style="text-align:right;font-family:sans-serif;font-size:10px" border="0" height="100%"><tr><td id="0%" style="vertical-align:top"><a style="color:#999999" href="#0%">0%</a></td></tr><tr><td id="5%" style="vertical-align:top"><a style="color:#999999" href="#5%"> </a></td></tr><tr><td id="10%" style="vertical-align:top"><a style="color:#999999" href="#10%">10%</a></td></tr><tr><td id="15%" style="vertical-align:top"><a style="color:#999999" href="#15%"> </a></td></tr><tr><td id="20%" style="vertical-align:top"><a style="color:#999999" href="#20%">20%</a></td></tr><tr><td id="25%" style="vertical-align:top"><a style="color:#999999" href="#25%"> </a></td></tr><tr><td id="30%" style="vertical-align:top"><a style="color:#999999" href="#30%">30%</a></td></tr><tr><td id="35%" style="vertical-align:top"><a style="color:#999999" href="#35%"> </a></td></tr><tr><td id="40%" style="vertical-align:top"><a style="color:#999999" href="#40%">40%</a></td></tr><tr><td id="45%" style="vertical-align:top"><a style="color:#999999" href="#45%"> </a></td></tr><tr><td id="50%" style="vertical-align:top"><a style="color:#999999" href="#50%">50%</a></td></tr><tr><td id="55%" style="vertical-align:top"><a style="color:#999999" href="#55%"> </a></td></tr><tr><td id="60%" style="vertical-align:top"><a style="color:#999999" href="#60%">60%</a></td></tr><tr><td id="65%" style="vertical-align:top"><a style="color:#999999" href="#65%"> </a></td></tr><tr><td id="70%" style="vertical-align:top"><a style="color:#999999" href="#70%">70%</a></td></tr><tr><td id="75%" style="vertical-align:top"><a style="color:#999999" href="#75%"> </a></td></tr><tr><td id="80%" style="vertical-align:top"><a style="color:#999999" href="#80%">80%</a></td></tr><tr><td id="85%" style="vertical-align:top"><a style="color:#999999" href="#85%"> </a></td></tr><tr><td id="90%" style="vertical-align:top"><a style="color:#999999" href="#90%">90%</a></td></tr><tr><td id="95%" style="vertical-align:top"><a style="color:#999999" href="#95%"> </a></td></tr><tr><td id="100%" style="vertical-align:bottom;height:12px"><a style="color:#999999" href="#100%">100%</a></td></tr></table></div><!--<![endif]--><script type="text/javascript">function showDivShare() {
  updateShareLinks();
  document.getElementById("SHARER"  ).style.display="none";
  document.getElementById("DIVSHARE").style.display="block";
  return false;
}
function updateShareLinks() {
  var shortlink = "http://archive.is/XvEff";
  var re = new RegExp(shortlink.replace(".", "\.") + "(#selection-[0-9.-]+)?");
  var adr = document.location.hash.match(/(selection-\d+\.\d+-\d+\.\d+)/);
  document.getElementById("SHARE_SHORTLINK").value = document.getElementById("SHARE_SHORTLINK").value.replace(re, adr ? shortlink + document.location.hash : shortlink);
  document.getElementById("SHARE_MARKDOWN" ).value = document.getElementById("SHARE_MARKDOWN" ).value.replace(re, adr ? shortlink + document.location.hash : shortlink);
  document.getElementById("SHARE_HTMLCODE" ).value = document.getElementById("SHARE_HTMLCODE" ).value.replace(re, adr ? shortlink + document.location.hash : shortlink);
  document.getElementById("SHARE_WIKICODE" ).value = document.getElementById("SHARE_WIKICODE" ).value.replace(re, adr ? shortlink + document.location.hash : shortlink);
}
function findXY(obj) {
  var cur = {x:0, y:0};
  while (obj && obj.offsetParent) {
    cur.x += obj.offsetLeft; // todo: + webkit-transform
    cur.y += obj.offsetTop; // todo: + webkit-transform
    obj = obj.offsetParent;
  }
  return cur;
}
function findXY2(obj, textpos) { // it could reset selection
  if (obj.nodeType==3) {
    var parent = obj.parentNode;
    var text = document.createTextNode(obj.data.substr(0, textpos));
    var artificial = document.createElement("SPAN");
    artificial.appendChild(document.createTextNode(obj.data.substr(textpos)));
    parent.insertBefore(text, obj);
    parent.replaceChild(artificial, obj);
    var y = findXY(artificial);
    parent.removeChild(text);
    parent.replaceChild(obj, artificial);
    return y;
  } else {
    return findXY(obj);
  }
}
var prevhash = "";
function scrollToHash() {
  if (document.location.hash.replace(/^#/, "")==prevhash.replace(/^#/, ""))
    return;
  prevhash = document.location.hash;
  if (document.location.hash.match(/#[0-9.]+%/)) {
    var p = parseFloat(document.location.hash.substring(1));
    if (0 < p && p < 100 /*&& p%5 != 0*/) {
      var content = document.getElementById("CONTENT")
      var y = findXY(content).y + (content.offsetHeight)*p/100;
      window.scrollTo(0, y-16);
    }
  }

  var adr = document.location.hash.match(/selection-(\d+)\.(\d+)-(\d+)\.(\d+)/);
  if (adr) {
    var pos=0,begin=null,end=null;
    function recur(e) {
      if (e.nodeType==1) pos = (pos&~1)+2;
      if (e.nodeType==3) pos = pos|1;
      if (pos==adr[1]) begin=[e, adr[2]];
      if (pos==adr[3]) end  =[e, adr[4]];
      for (var i=0; i<e.childNodes.length; i++)
        recur(e.childNodes[i]);
      if (e.childNodes.length>0 && e.lastChild.nodeType==3)
        pos = (pos&~1)+2;
    }
    var content = document.getElementById("CONTENT");
    recur(content.childNodes[content.childNodes[0].nodeType==3 ? 1 : 0]);
    if (begin!=null && end!=null) {
      window.scrollTo(0, findXY2(begin[0], begin[1]).y-8);

      if (window.getSelection) {
        var sel = window.getSelection();
        sel.removeAllRanges();
        var range = document.createRange();
        range.setStart(begin[0], begin[1]);
        range.setEnd  (  end[0],   end[1]);
        sel.addRange(range);
      } else if (document.selection) { // IE
      }
    }
  }
}
window.onhashchange = scrollToHash;
var initScrollToHashDone = false;
function initScrollToHash() {
  if (!initScrollToHashDone) {
    initScrollToHashDone = true;
    scrollToHash();
  }
}
window.onload = initScrollToHash;
setTimeout(initScrollToHash, 500); /* onload can be delayed by counter code */

//document.onselectionchange = /* only webkit has working document.onselectionchange */
document.onmousedown = document.onmouseup = function(e) {
  var SHARER = document.getElementById("SHARER");
  var newhash = "";
  if (window.getSelection) {
    var sel=window.getSelection();
    if (!sel.isCollapsed) {
      var pos=0,begin=[0,0],end=[0,0];
      var range=sel.getRangeAt(0);
      function recur(e) {
        if (e.nodeType==1) pos = (pos&~1)+2;
        if (e.nodeType==3) pos = pos|1;
        if (range.startContainer===e) begin=[pos, range.startOffset];
        if (range.endContainer  ===e) end  =[pos, range.endOffset  ];
        for (var i=0; i<e.childNodes.length; i++)
          recur(e.childNodes[i]);
        if (e.childNodes.length>0 && e.lastChild.nodeType==3)
          pos = (pos&~1)+2;
      }

      var content = document.getElementById("CONTENT");
      recur(content.childNodes[content.childNodes[0].nodeType==3 ? 1 : 0]);
      if (begin[0]>0 && end[0]>0) {
        newhash = "selection-"+begin[0]+"."+begin[1]+"-"+end[0]+"."+end[1];
      }
    }
  } else if (document.selection) { // IE
  }

  try {
    var oldhash = location.hash.replace(/^#/, "");
    if (oldhash != newhash) {
      prevhash = newhash; /* avoid firing window.onhashchange and scrolling */
      if (history.replaceState) {
        history.replaceState('', document.title, newhash.length>0 ? '#'+newhash : window.location.pathname);
      } else {
        if (newhash.length>0) location.hash = newhash;
      }
    }
  } catch(e) {
  }

  if (newhash == "") {
    SHARER.style.display="none";
  }
};
</script></div></div><script type="text/javascript">
var _tmr = window._tmr || (window._tmr = []);
_tmr.push({id: ""+63*44843, type: "pageView", start: (new Date()).getTime()});
(function (d, w, id) {
  if (d.getElementById(id)) return;
  var ts = d.createElement("script"); ts.type = "text/javascript"; ts.async = true; ts.id = id;
  ts.src = (d.location.protocol == "https:" ? "https:" : "http:") + "//top-fwz1.mail.ru/js/code.js";
  var f = function () {var s = d.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ts, s);};
  if (w.opera == "[object Opera]") { d.addEventListener("DOMContentLoaded", f, false); } else { f(); }
})(document, window, "topmailru-code");
document.cookie="_ga=GA1.2.661111166."+Math.floor((new Date()).getTime()/1000)+";expires="+(new Date((new Date()).getTime()+2*60*60*1000)).toUTCString()+";path=/";
</script><noscript><div style="position:absolute;left:-10000px;">
<img src="//top-fwz1.mail.ru/counter?id=2825109;js=na" style="border:0;" height="1" width="1"/>
</div></noscript>
<img width="1" height="1" src="https://1.1.1.1.us.SEF2-x.215506745.pixel.archive.is/pixel.gif"/><div style="padding:200px 0;min-width:1028px;background-color:#EEEEEE"></div></center></body></html>`: "",
	}

	i := 0
	for content, expected := range testCases {
		err := checkCrawlResult([]byte(content))

		if err != nil {
			if expected != "" && !strings.Contains(err.Error(), expected) {
				t.Errorf("[i=%v] Expected checkCrawlResult to produce an error containing %q, but actual=%q", i, expected, err)
			}
			if expected == "" {
				t.Errorf("[i=%v] Expected checkCrawlResult to succeed but received error: %s", i, err)
			}
		} else if expected != "" {
			t.Errorf("[i=%v] Expected checkCrawlResult to produce an error containing %q, but err was nil", i, expected)
		}
		if expected == "" && err != nil {
		}
		i += 1
	}
}
