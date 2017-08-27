package assets

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io"
	"sort"
)

// AssetNames returns a list of all assets
func AssetNames() []string {
	an := make([]string, len(data))
	i := 0
	for k := range data {
		an[i] = k
		i++
	}
	sort.Strings(an)
	return an
}

// Get returns an asset by name
func Get(an string) ([]byte, bool) {
	if d, ok := data[an]; ok {
		return d, true
	}
	return nil, false
}

// MustGet returns an asset by name or explodes
func MustGet(an string) []byte {
	if r, ok := Get(an); ok {
		return r
	}
	panic(errors.New("could not find asset: " + an))
}

func decompress(s string) []byte {
	b, _ := base64.StdEncoding.DecodeString(s)
	r, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	defer r.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.Bytes()
}

var data = make(map[string][]byte, 69)

func init() {
	data["core/00_builtins.lisp"] = decompress("H4sIAAAAAAAA/6SX4ZLjJgzHv+cpSGY67bXj6fTr5rZ5FgXkmFsMDsi5TZ++A7ZzMcJnZ24/3Uk//ZGEYuB4PB5F6Hoi9EI6j2/i3GtD2obd7g+FdTX+VyishXhTTlYQApI4KKwP4i10KDWYqna+FeR7/DIP07WYh+n6INbDDNI8zCBtWE25bDXltqzmEdQsLBoOGYQ3MDMoGg5fdrvjUbRIoIAg69l3TU0VfbO4hzVfIZHp7xlfIk+MbCAUdfcPfIFOJYD0rrDpyS4Ez2vybdiWxOFnB1b9k2ReCBCvBlRgzHrAtXeE4sffc2XJdxBvQ+kxeF0v3C3BZ/WQfQoupXpiWzRa41ZYaDF0IDHfjjQ5NvBpsmHDLtjAS42Bix2OuZwhaCk6j0pLIJYRXnNFvLLxS9DPGavNKftOhMpqw7QSuM594P2784q1+Yc9FhfIa3vJawrks3wD+XyBQH6/BeJVFbh9AotcShOvPVo+DrX2gWZhyZKrewzZxzRackq6OB/PVLQUqG+M+pZTBq3IP9+WbTk1+TBSk0OXeA7MoAvS2JbFoQx4PfFRjy0tDOcDXmcN2rJuobr9A15ntVX4iYqPyuhg2o+ART72B0JwUgPpG+tQcuXJJWO+VjKeeBlFeD/RW+AWuq5U9GDnZ9iIL9GxYqMD5aVGWz6Mmo9/tPFcSuR+QMtkzOKGkpzP8xisWSaDMV9hsPJsyvR+wpfodKT0LXots5z+GpYQB1AqCBAx1NURPqNnv/1qokN/Jg+S1kP+nELa3pDujMb1mL+nGKVvWm0I+G0K6LxTvcQgqEHhsYX4e/Cidl6AGNQU15qLvU9iskH5wRdPanjtwWi6s6143xCs7VL4v9na1ABxDW2lRwgY/yVuYHpkOu9bdJTDYH8noXDQW5T7uimtSWY5ra8vpTWV+SyXfzbr6cP01PFBJ3W6c0HHr18ktS3t1yDxCwrVmoLFC/w8h0HiFxQsWNaHsbepcUIHYR09VPmdKgm8Eh8/KXVvJWnHHorSuNB75Afl6HjxYm2gPSso3NIHx5bnHXSduZcu+snBz6UPrKYilq/xKZbdL4vW5/xYROaMnZXOyt57tPK+K+R2cVteRLIBW9iDBtj1o/Ou1QEzcrQuwPxwWuD3j4BFPp3c8N998Z6bqo5EvJStV147XyHIZvlhk19pJVBpPAZPXlGtDaEv8YOncM9ZeGe20PG7ejzBSvDgyXmCDyyLR09OK++6Mh09xR+CB3sZn4IdWC1FfJJ7lO6GvrRNo2u97UltnPc07CRIt5jtzeCpoufL7v8AAAD///QOnZCtEgAA")
	data["core/01_branching.lisp"] = decompress("H4sIAAAAAAAA/5yRzWrrMBCF936KQy4EexG43cabvEcwVJHGsUDVBEl2CaV59qKJkygJhbYC/30jn/kktW3bIh7GlChAc6A1dkF5PVi/r6raUP+mdGB4ThXwsTassXC8t1o5d4T1E4UUkQZCCmMarKcYwb2QQ+DJGjKYlBtp8VkB20m5rgJea9vjNCmHXrkoP1PTlB3fB/KXlisVIyUsMjvHJIoJS+zYHG9xAmvDOG1yoYG37jl0VazlN8HeujL8PlizN4+hmUlove1ERV61U2OkDufnmS3nr5hbXprWM8P/RqiMG30paD2f4Dq3xGlTB9kJuc/zGxmlsno2VoWwnMjPjR0lbJU3/+7Eu9kxL0iKV1NViF4VZc6jJ4dHTQ5/3FaR5PC9Y67l66rJ4dlSBL8CAAD//4iShwk5AwAA")
	data["core/02_functions.lisp"] = decompress("H4sIAAAAAAAA/3ySwW7jIBCG736KXz5k4ZBVnNxsaTfvEVkKgXHWWgMR4F6q+tkrsNOQusocrNEw38/Mb5qmaeBvYwjkIK2jGt1oZOit8UXBFHVaSGchvCcXtsJdfQGwUwvTDzylchCjpxbsJkwvUWvyXlwJZcZgbvLQow+4EG6id6RKPkts7udtAQBnJq1RKQUmdmd3HGwZtc61pyNz5APm79LNk3SMmgZPWOL7kA/1isfId+5MAbzXysptvC6g7Ez5UQCnDTrrdJqWDb0P+HWfS4v/tJWD9aOjdH++ChNGgf0BG8jMCjwt9UYyWPcXLNWw45w/8hdstWarjK1esfs1u8/Y/Sv2sGYPGXt4dv7U8i8n/MOpQeiLEovks/GK1tbH2my+EZqQ/YFzBDGlMuvMkk3Hn5R/R4Eg3JUCNIV/VmGD+IhmIRbr03I+zQ0c0zF2cF58BgAA//8+3X8ULQMAAA==")
	data["core/03_concurrency.lisp"] = decompress("H4sIAAAAAAAA/3RPTWrzMBDd6xSPGL5vBM0F5E3vYQJ15YlranvSkbQIpT17saxiJ6HaaHhv5v3UdV0jXFKMrPCi7OBl9kmVZ381hjo+T61XQS8G+HSd+GMbAkccejl8GaD5h1fpricD0DiEiP9Fz03tOx/9KCEpo1kWAPIyh7uVXrKCtfbGkGfWNvKDbcHvzV9o5IjGv7VzBVo++5Qt81ticAVyeUDe2vM8DRHk8reSJW8vWIWVQxpjBeoE38858Gm7B60OdgeVC7ti5AJ/FO3bpucUkz70XNE/Wl5UpiEXKtM+7kZuUUuIX8pa8xMAAP//D2AMr/sBAAA=")
	data["core/04_lazy.lisp"] = decompress("H4sIAAAAAAAA/4RU7W7bLBj976s48iu9w9Istf2ZaGrvw7IUih+naBgcIGuzab72CfBH6jlL/iTmfPBwOPF+v9/D9WfvyUIYSzso/vMCR6czaUEuy1hDbceFNfCm5M4ZkQG/do0R4Yk88mk9/50B1f9B7OoMODDe9+qCCIIJowX3GF4CXhTFylpJ5zecw/It44Dd9f1Bwhu74ZyAW94J/be75fpIa+e4GF1ZVeP6c2Bj0ruOf6cyEvEwxb+TusVjUURhxz/qu8Kh4x+LQmosqk3FEDhbIjhPfX1PNATWKoFQltLRaR3CtD6l+2qaS0iXxSv78mkfoYw7W0IVCIiJuxVlsos+qxHe30iXikJ3qlepG6mPDlc7homsL7k9uuSfbvYZE7kAkI8P07V3Z+fxSnDnvleSmjxJv+EJTJFetMW2VBjtudR4AinqSHuXx8QVeVStsR3YPOtD8TWbbtpTaPQMPRYpk0MSeur6/zAEUj1KWDg+IjCbJPIQt4lIjeFlTG7KTsOHYN/fpAoNrnpLDYRRajuz+J94RmAVyMOXFNzTrZwcnZ6jW8w2/CDhpdF/81MqS4vGE5XxCG6eaDkqi5OyVlrn4WKPRzj2ZgYwl3k5KJLWUpJexZGq1Mb3ROXoVNJHb69rdGBH0mS5pzRka2xJXLxhWNiMOunBGvMp7j8BAAD//40EpERjBQAA")
	data["core/05_io.lisp"] = decompress("H4sIAAAAAAAA/9RU247aOhR9z1esw+iMbNR0mHkETZn/iCLVJDtgybGD7YBGVfn2ynYuQK/qW/MQkbDXZS8v2Gw2G7iu954sKmNpDflksozV1Gh0Nm9Fl5+lP+Raqgz4su6sPAlP8LanrxlQNL2u4OhYZgBrRZcBAGs0ipNQJZhswLRUW5yE4uEGFiHhkfM47ejI+SyK4hGNsa2LlIo8CkdHsDs7cN7+lwZ5mVSDlqPjNjLGVwD7OC64Xjpfm94vsT5b6Qmskdb5JM8H38bmJKoDClLUglkaBsrf0y1dJypaIiA5v95I364kuk69h0WT+ezXrJrOSmpaXhNK7f8wpQ+Db4S4/o20pPbqJ4mFvf82NCEt1fmJKm/s9sddDo2NcroGGyZTb9kr2P8hZp2eXzhWk+lWVNYgJm460ih2UtdS7x0esTP1e+J0jqzPhd27FN6tIYwYjsXMNBG1vfPYEcSwBxJsEXOojK4T52uyOHOtxoP9zMasaoPLW7A1HOKn70AvMyj26sImHyuOq6dnPtcrXpUyjh7A9uRxi1rHr6ClusdYcr3yD5gMXgV5eUudSvfJIS/HFcqJi50PpIembAcjHGz4wLM7ubvDi9XK694KL43OWxf/2oxtg8AQg/PCBp9Vby1pn3vZ0sw77XEJsPk16foeM/z+xqazJ7A8zSUJjudVvDgWrVskjcl39i0AAP//SaHTXLUFAAA=")
	data["docstring/and.md"] = decompress("H4sIAAAAAAAA/0SOQa7CMAwF9z3Fk/4inwp6BxasOYPVOk2kNK5sl8LtUSgSO+tpPJo//FOdEEWXvj9hZW2ngWBJ1C9j1nHLnuuMQHUK3e1BZSNngyfGAUeVBYWjwwWa5+QDrgYTqSCDVAb/3gSEEKkYvwLazGfsuRQo+6YVnsiPfcDdE+uejZH9gFaVkXlqmq+ztbWWyk//BA3dOwAA///2Hivp1gAAAA==")
	data["docstring/apply.md"] = decompress("H4sIAAAAAAAA/1TPTUrAMBAF4P2c4kEXJggFFdy76A3cl7GZ2kCaxPzUentJgwWXM3zvDTNAcYzuB2v1C7J8abTZSganz7qLLxklgC9QbPA0HewqF8komyCmcFgjpmWr+EXA3twdtmQ0LlfJP//XBxVS9wt/ONHgHrvPj0TDgDeP6eQ9OiEClJEVJx7UE57xonVb9T8ecWqi981mSPf4ts4hSanJY36dR/oNAAD//8CDF8f1AAAA")
	data["docstring/assoc.md"] = decompress("H4sIAAAAAAAA/2zQz0oDQQwG8Ps8xcf2YosWwZuI4EF8AcFzOpvpBmczS5Kt7duL3eIf8Jjw4/tIVrgi95bx8M4nHKjO/LjZrJGNKdhBUP7AmQiFHBgeNueYjdOb1HqB/7gdOfdoihgYk7WD9NzjuwUTifk1msE4ZlOQgscpTn9SpEAbSrPRQfYTtMXrII4yaw5pCvEv10spbKyBGGhp/h1WJdiowk8adAQfM09nGpBAJsWOEeeDeoiCUMQ8bnIldxTyQZpuU1qt8KR4PtI4VU4Jlx8mALhXGhmdT3MEW7fsaM/A7fZumSrpHt1L69bpMwAA///9TMgsfwEAAA==")
	data["docstring/chan.md"] = decompress("H4sIAAAAAAAA/2yTQW/bMAyF7/4VD+3FDtr8gNyKoafdht2GAWZkOhYmSylFJc1+/UA5SetiucQQn8j3PkmPaN1EsYMTJuUMQol+8MJOfYoUYOXIoXm5fcGbaiAlZJXitAhDJ1IrlMwDNOHAkYWUQQj094LMb4WjY6QRJwqF8xY/J4ZwLkGXlhPl6XmmI1yK2Wf18WByiuh59tpjLLG6egKhdyFlXq3FAXQftMUPVvF8si4UwYFnjopR0gyd+MPRTBds9iG5P5snnMnXuWOSqor8rothS7VnmBPlAUnumhuXRVB9DVu8mrAO/9hPd+3ZhwAKOaFORonqQ+22L+PIYkTGUPLEA8jgXEGlEV6NbIVzYlll2TbN4yO+XWd850tums0u89sG9lvFXtzezmloNjuLVoWGq8aUO14bbNvHJDP61spLrK5vNrua2bbSxwZNC4oVor6ta11fjb5EvL7TfAzcNEAbWPHLTdcb+bsxK+0h1X+g3cJNqCbxMKb00P2vsCf5Wlgmds3ST9PziZ0mQWtg4KbOaldmGcJH4Ww3xXyHdH4OfOKAEgeWo4/Rx0O+0ehv+D7fw/Pk3WTn5yPmlBWOMueqPwqPLMIDznRZPZPVI7m+DV7YgPbpxHCphMGOTPgsdjhxV6l9CnTrtiT9Sqr9RKjrmn8BAAD//w/0PTP6AwAA")
	data["docstring/concat.md"] = decompress("H4sIAAAAAAAA/1SOzWrDMBCE7/sUAz7UplDo772UvkFuIRghbyKBIjlayY7z9EEymOQ6336z06DVwWuVIHx57eDUzboFa8ZeJZZCMnvNQn+Ra6LK3bIBzCYIFymxT7CCZBiRJbuEcHxos/5U2aRcZimMlTYYY5jswMPW+EbUNPj1+L+q8+iYCNvQ/Ts+8HnAS/uFb/x0HdHOWMFsnUPklKOvP54n9m3VsDo93QMAAP//uAS+0/kAAAA=")
	data["docstring/cond.md"] = decompress("H4sIAAAAAAAA/7SSMW/bMBCFd/6KB3uIFQSGMmRIEKTo0AKd2z0+U2eTCH0UyFMl/fvipMQtCnSsJlF47+m+e9xi57N0eO4Ld9DA8nJ7C06VPzXouZxyuVSYJGrMQgnHQuJDlLP7mguYfMDhD/cBPtFQ+c6OsO/RkzLGmBKODP5JaSDl7g4kHeIJURErtAwaZuwkK06ULMBeJaZmSbqx9BvYPCa/xiwphXUoYplZA5cxVl5MwpO+z2OmvmTPtXK3d267xWfBl4kufWLngF3HJ0x4fGyWkyE7ANg9Y8JD2wCbkSoS1woNJHhoN6vgBRPu27ZZBefCpFxWzX37Lvr9LKIoOLKOzLJpnPsm0BArPBn269+K1+vyPjj3+J7iOWiaccmFQVw1sEZPKc3oE1ONcsaYh9SZTzNsBSR4smrxxvOYS/f0P0jXP/yD9EcwCKrZiElX7DGXt2r9HNkvVVl1hyXn8DHqtfBqMOtl2btfAQAA//+f2055vgIAAA==")
	data["docstring/conj.md"] = decompress("H4sIAAAAAAAA/2SQvW7jMBCEez7FAG7OuIOAy3/rIm8QIEVgWCtyZTGmdhVyZcV5+kCR7SYt55vhzqzwx6u8o/AHWs393zUohAJO3LNYgSloVkcWz27zW5vtGoWaxFeuwksXCxru6Bg1Y4opIcS25YzAA0uIsocKrOM5wGc2hp0GrrBBisUWy5B/4H8gHNnbJYmG5XXqYmKQgEpRH8nikdHTgQtEsR8pkxhzATU6GjQHzlH25+vaUbxFlSVT1DBpPmCK1iHR1+napqCMvgMVqHDBkDWMngOaE+qehhqaUbcxGee6cm61wkbw/En9kNg5nBd++48b3OJui3s84BFPa+de558z25iXKc4ldxf2gm53lfsOAAD//0Od9nWqAQAA")
	data["docstring/cons.md"] = decompress("H4sIAAAAAAAA/3ySP2/cMAzFd3+KB2SoLQQHJP2zd+jWsUDHmJF4NgEd5ZLy/fn2hexchqLtyEc+8kdKD+hjUceZ8spw/jUgltOrKDvoTb1InUEtubJG7n5usaJYYuP0nniEr3EGNWcWryiGM8da7BF1Zhj7miuk5ZUvf9TsY1rdPnUxXlgTJ9SyycVkEqV8wB1g1f8gKMi9RKEqZwaZ0W3DMG4EWjCtZKSVG9hElkQnLMWlStFD1/2YGUonxtgONO7YZ/YqE6MccZlZ8V18gZyWzCfWSs3qCM0Qqq2xcnrf1NuqkXN2vN6wkFgbSBgj2Yg+xFC0sra647ZvoJCSsTsWstrUYIEn8co23B9ljOkf7hQ42ob1N/+h6x4e8FXx7UoNv+uAPvERV3zoP+ITPuPLMNzF29svecZ1F7foCbcBaJcSB+99cJGcYVxXU7z0T3jGvdvLofsdAAD//yLnvbNuAgAA")
	data["docstring/def.md"] = decompress("H4sIAAAAAAAA/2SSzYrbQBCE73qKAh/WXjaGbJ5gDzkEcgzkEALbkkqegZmR0t1jW3n6MLLxHnJpRP98Vd2aHfYjJxTJxDRrPkDM4qkYZEvaIgNxFo3SJ3Y/Y0r3DsgjjVjggRiqKot/DL7gEuIQEA1P1ahP6FeMnKQmP+ItpQfCIErEnKtvRCkjlFaTN7gUUHVW9IzlBJVoHBGnVhB35sWbRpaR8BnKT3ePHpiP+BGioWeQc5y1NY5xmrhZ9SAFeTbH92iLvUCsDa24tE1PLFRJacUkMcFiYvG0NktWh4BBjHbsut0ObwVfr5KXxK7D7ajXDgD2WZbtA9hPBb/W31jx/IzXwz1r/PO5hdcWvhwOXbf55Y12MzIoxQlBkr8rsiy4BCpBGQLOkioxT9s/8KAkFp3PceTYkJVloG1rz7VPHFGXuUBbxfyIb37T8MCCx93EP7R83tD/vwe8X9+P3b8AAAD//70yXNVEAgAA")
	data["docstring/defmacro.md"] = decompress("H4sIAAAAAAAA/1SQT2/CMAzF7/0UTyCtgU1oXCdNiMM+ww6Ig0ncNVKadLE72LefEv4Meqhe/LPfszyHcdwNZHNCpIExsNIGu6KXyz26lIfnBUjEf0UBITM5zqgTzacPAY5toMygcxE+QnuGnXLmqNVVRrL8gmPvbQ8vaCfh3OLwC8cdTUFX2F6nBXwaKTp2hRenS6SPSLkITaCgRfQMSVO2DJsctwJHWlYcMwtHJfUp4sBdygyv1fuHwkTKbtU08zm2ER8nGsbATYO7U9gUXQMAM1/gwFGlxrWFtPUss9qwe4INNAnLvr7NsecII/y9uYJFBYDxHcw71jCB4w1eKWAuJbze1YIXRSvjpMr5zXc38tD/r9eLh44U5bwzTGbRy/+WXb/mLwAA//982CfHBQIAAA==")
	data["docstring/defn.md"] = decompress("H4sIAAAAAAAA/2yQzWobMRSF93qKQ7zIDLFN3GWglC7aJyh0EQK5lq48l0pXjn5IndJ3L1LS2IvMYqRzf4++FSbHXqEUGZErfcE95UN5gE853sygUuSgBTRKypEswze1VZKanxICHNtAmUHvcYiiLgzbcmat5841nhexC6TguhXO19if4NhTC3VrzGqFr4pvvykeAxuDN29e9gBg+u/PnUsWV5lry92VggPHviT5sdPLPilZKyj81FgtX/0dnffyMM7JJnXjBkyfIbidcXupdzN2l/rTWd9xKNzjN5i6q2kz8vOF2s39M+bHIgX8+hQcOXecBZlty6Ujepa6QBMqSYClEJCOVaK8UCe4Bqnr6cgxvcW2+J4yCDHl94mkdVSWSvbXxmdhdeH0AQQcWDmPOf9drVGYBzOXbOsMX9PJ4zHQy2lT+Olxa/4FAAD//2ZYi2siAgAA")
	data["docstring/do.md"] = decompress("H4sIAAAAAAAA/zzOParDMBDE8V6n+IMbSTweSZkmkCJnSC2Sdbwgy0YfQccPdsDV7vCbYgbsa2Fc8uy9Qz4htlClMLdYdY2yUzEPjfFQJDynHdBEbTn9kWW7mt7USRg1haOvSyIUtBaylBbrvzHDwC1x72FeoxgDVkfslc7FbYsMgF1zorvf7z2d82lLzplvAAAA//+liQqetwAAAA==")
	data["docstring/drop.md"] = decompress("H4sIAAAAAAAA/1SQT0vEMBTE7/kUA3vYBqWg+0evHjx4FzyILDGZ0kKa1ORVWz+9ZLvq7vXN8Jt5s0LlUhxg4xgEmR9XGuWQIS3RdCkL6NkzSEZsimNksMzqpfMeiTKmAANvvuc/EdIaASfrR8dz0voYs74gFnVI8bNzdP94VDYGa4TBCJ2u8bQ4cxyT5ZnPJCK3MQlTiQ1HVypyFrrlr2uYAPaDzPBdFnyV6u88taerlVqt8BDwOJl+8FQKqBwbTFhXN7jFBlutf48zXnfY4w73b8ut7LfBhFkr9dx2GVwwS85polLrcqRDtcUOe32o1U8AAAD//6EiwBaGAQAA")
	data["docstring/eq.md"] = decompress("H4sIAAAAAAAA/3SQsW4jMQxEe33FnF3Yxhn7D1dcESBAypQGrR1ZQmTKFrlx8vfBeh0gTRoSGIAzj7PGllek1s/38XcHp7mhJAiMjpbwLnWiQTpRRqqXKBXe4JlIpZuH11IrOn3qikOSajxADNaazrs4qLFN6uwGWeI8i6MYtPkvtgNePLPfihG3nwneJx6GENZr/FP8/5DzpTIEANiOTMhYZdbaVrtF4/VbQN6F8FzeiHMzx6VzLFGctofnYkiTRi9NEUUh1RqOhPIkzhHHz/ngQh2Lnu6Ymz8bxCxdorMPeNLFJYpx/8BcyI98wHOcq51/llrvRSzF8jpJHcJXAAAA//+7FebRjwEAAA==")
	data["docstring/eval.md"] = decompress("H4sIAAAAAAAA/1JW0EgtS8xRSMsvytVUSK0oSMxLKVZIzEtRAAmXJpakFiskgmW5AAEAAP//Q1NDtisAAAA=")
	data["docstring/filter.md"] = decompress("H4sIAAAAAAAA/1yQzU7rMBCF936KI2UT615VasuSDUIs2LOrqsqkYzLSMA7+aROeHjkRgbKz5xx932gatJ4lU4Qv2iHRxz8LcZ8sE5Yg1WEh7SiZx0guU4KrlWkNcO1DInRBM2kGJ+SeECkVyQgebhhkYn2bx0MMFz7TeRZmDooc5uDipFCq/Zvaat/g2f8B11+Fc+dmUjXHkvsJrYYM7yTRf9Snstha1x8VriyCVwJrJ6WqWH/x677f7o0xTYMHxdPo3gchY7DerfWKw3hEe48Re2tx2GKHPe6O1piXntPiiZRLXPi3tzu1W+zsyXwFAAD//1JglZSKAQAA")
	data["docstring/first.md"] = decompress("H4sIAAAAAAAA/3yQMU+EQBSE+/0Vk1whJIZCL0YqY2Fhb3/AMisvPnZhd/Hw35sDo7nGdibzvTdzQOEkpozEuURkXqJPyAOxy1SO9BnBbWLivNBbmrdBEtzibZbgcRbVn/A/2YlWnLD/pdwiRHhRyDUdksBxyl8VXjPOYdEeHdHR04mVVpED7ED7ARciWny2Kn9YLEn8O5oicX7aijXo6EIkbKu6eduHzeV+E5lyUxlzOODZ42Vtx0lpDFD0dFhxU9Q1Ho64v8PxsSwvxt5vLY3BNgT30NUOp7o+VeY7AAD//0ew0thgAQAA")
	data["docstring/fn.md"] = decompress("H4sIAAAAAAAA/2yQza7aMBBG936KT5dFEypQSSnqDrHoG1TqIorQ4EzAkn+iGZvC21c47Hq39tHRnG+FZoqIFPiIwJmO6EmuOmBKEr62sMKUWUERFFN8hlQUU4k2uxTNH+f9G/mcQL5RRqAnLoyZVHkESSpxhIsgTE40b6wnVQSKkWVrzGqFU8SvB4XZszFAM/KEMZWLZwMAjeeMPhSPbqgPqBUfC6K4ky+sH+gfA5r1Gg+E4tu2bassp82dbU6yuALNbze+NDt0+I49fuBQ8d83p+DlEvx91QrnIq8uxmLBue+wxwE/sfuGXTect8ac/ltCMUsai2UQrE9ahJdxbJod6yIkcVQL6s+N7owLc4TwxMLR8ohJUqiwFqlDuniF2jTz1vwLAAD//+LBDUrNAQAA")
	data["docstring/future.md"] = decompress("H4sIAAAAAAAA/1yRUZLbIBBE/zlFl/0jqxLfIR+5QGoPIAytFRU8uGDQrm6fAinOOnzOMP1mus8Y5qo1E3PK93G8gKuN1SoLLAoVae6tAls2cUtOkmqJm/lFrVnar0MgCHTZddpQq4vTkOSKt4V45LQGT3+ofYQYceMT59t8Iz5stkrokml9E+InXW0632DFw8oGZ2Ms0LQDD8wuOd5icr9HVNEQnwsVLHYlbqTApfsjUhm3f/CrMeczfgh+ftrWNQYYPOd221+HDAAMmr6vdJoyhncK2669AQy8B8Up058ur6WUrbzz/+rGGNPH6dJe581VL8a8LaGA+xrNNF8dv7jcTqfHNFedoItVuCQrs5Z+bGapUUsPQF4iQzcGQTTBYr9hT2bq5KlLH7n0r18s3Nn5GXnzjVfzJwAA//872/RsPgIAAA==")
	data["docstring/generate.md"] = decompress("H4sIAAAAAAAA/1SSzW7bMBCE73qKgX2x0sTv0ENeoAh6KYp4Ta4sotSuyh8n6tMXJCXbOYrkzjczqz0OFxYOlBiDhulbj+07ghD5b2YxDIqLmDGoaI5+6V6v5HN9k0ZGnNm4wbGtEhFO6uhMVTaNgclCB/Anm5ycyhE/OOUgXxBppIQP5z1caob0ygEkS5ktnAKtSEqgwODJpcT2iJ/t4uEMOTq5gODVkPcLotG5GMxiioNNshjG6VCmmnx/OuJtvVgVq2u2oGLX6EvQnJzwMyYmKZRmCAWEpNVCC/J09mr+PCFLcr7yjIbAcVaxZZA9TywJLiJwVH9li/NSKRLzxGGzuZV07Lr9Ht8Fr580zZ67DjhYHmDUa4j3ZXYA0GLtAttd/3iggeTCX88W9l4/dn3fV82kL1c2ScOq3Hfd2+giuHFbPIrRXaQa9PRvua8y1O2WKOXu9oNtBbVdBkdnzzg1wFr7A7c8rpha4q1DSvdZFzHkst21L/sMElv0ZfNQWU3x/Vet4pb/Fvr3+7H7HwAA//8ocJVqCQMAAA==")
	data["docstring/get.md"] = decompress("H4sIAAAAAAAA/zyQP0/DMBDFd3+Kp3SglVAnWLogkBhYEXvlxC+NRWtHd5f+EeK7I5uko++ef/e7W2F9oKHPcsI3bwjs/XS0lw2EJpFnKjzO/jgR7a1E3CdtkqSwgXPjEm2ICR5dTuZjosAGb4gKr5q76I2hpuonHdnFPjIU3BYffa3W6ZmKlA28RrWFW7p38mN5JjDaUMdwUZ5dohbzSRIDssAnUCRLrfuoDFvnViu8Jrxf/Wk80jlgHdhDcksx/OySPxHNW24b7PyBeHr+3ZRQOdQc2vkQhKpoLpLTodk49zVEBf+Zs4JiP/f3aNn5SVmNl6PEM+vK6cGWBYvwHT5KHil227q/AAAA//8Cb+cmpwEAAA==")
	data["docstring/go.md"] = decompress("H4sIAAAAAAAA/zyNTY7rIBgE95yin72BLHKHLN4N5gJfTNugwXwWPxl8+5EnUralqu4ZdlOsWvbbzUHqmZdQNGuv6QRfkro0VgieSZdv8xWIo+grevq/quInpoQnP7JHzBBUHlKkES0Uioeu4ODSW9R8xyOfKKw9tZg3XCUvowW+jz6rPtZFiqe/GzPPeGT8H7IficYA1nPFgD2K7rHSuYttCjswBaak0xvVVmCHw4RT+7/Jmd8AAAD//0iWKE72AAAA")
	data["docstring/has-meta.md"] = decompress("H4sIAAAAAAAA/5SRsY7bMBBEe33F2ClsOYmB5AdSB0jpLgiElbgyiVBLgbu0nb8PKPnOh4Oba7Uzo3nDT9hPbPQDY8rT5xbGaoqrZ/OcYZ4x53QJjt2iUHi6MKrFkVHzcwTFeD/xhWIhY1gCiSQjY4f6kfVLzRKYD4qxyGAhCa4hRmS2kgWd5cLdESfPGENWgyT5+oi5a7VEQ5Cl2WvOGhHkjG6kqNwdmwbYOx4FLpU+Mn7f/mB/OOCG721bjyv1em2b5lSL8Y2mOfKTXuh5oKKPfyoGEvT8AD02za/wlzElNcyZXRjIVvC3zNVGUVP1Cp8Xtv5fNcwsrjJUtN1mh8FTpsE411GCYmIShXmyZfM0PnueqSzLLZJ0fWkXkqyTbFbsrecY0xbfPkCuloOcFSQOUqae87LBzt6t8D8AAP//HDug1VICAAA=")
	data["docstring/if.md"] = decompress("H4sIAAAAAAAA/3SPvU7DMBRGdz/FUTPEliIkKH8TiIGBx3Ca68aS40S2Q9O3R04HWNhs3e+c+90G7R1LkoEySkRClnfDIsnNacpkPy1B6JONp9HHs/pyNYh827DaIsPO+pMtgs+UtJbxio5zwdmQpaM+ow+m27m2bmmp8pr/1dg4kKSsKcrQMZdR0sVnuUG11Q3q8A4brx0XHwK9/KO4U6pp+Ih8brZeoBToQRwbrb7ngSOPPPHMC6/G7FPv0G/oIJHNcDQK4LDVmr0/H/5882RDOBj1EwAA//8lsX8FPQEAAA==")
	data["docstring/is-assoc.md"] = decompress("H4sIAAAAAAAA/4yRsc7UMBCEez/F/Lni58QRCehoEAUFEuV1CEUbZxNbOHbk3eQ4oXt3ZOc4AaKgtLwzO/vNAS9IJNn3GFOeXx6hLCq4OFbHGeoYS06bH3ioEwLKjCrxpH5jMZ9GUAj3X94orKQMTaC4D56KTYQ6LxjXaNWniIsPAZl1zRGd5pW7FmfHGH0WRUzxVdX+mpM1KHysgR4eu9zHCd1IQbhrjTkc8CHi43eal8DG4HHfj3eRZkbT+xCa2+O58DRdmxu+vMYbvP16NOZccvJu8GfMfQl6trQK1yzqfN7JwAsIG1tNuTXms//GmJMolsyDt6Qsp78Y2EIoSELPiDyR8oD+WgQLx6HcVVY8Pz3DOspklXOB5AUzUxSoI63s0/ivpua1ktRi/3tjbcXydOfSOA4hNf8FYO/J/AwAAP//ysZCyzUCAAA=")
	data["docstring/is-indexed.md"] = decompress("H4sIAAAAAAAA/3yST28TMRDF7/4Ur80hjSiR+E9PiAMHJI69IRRNvG+zFs548cym5dsj75aCSul95vn9fuMVLpJ2vGX3AX2px2cbOM0NNwN9YIUPxFjLKXXs5gmDVOJuCcYfEzXSwucekvPdCE+SJ3HCCwQnyenPKHwQRRTFnpAYacYO+59L5mV7UeFDMvSTRk9FcZNyRqVPVbHzOnG3xfVA9KmaQ4s+f9jn94pN2ZF0xriPW5KSHrDrJRt32xBWK3xUfLqV45gZAv7ysr54gZd4hdcbfH2Dt3iH9982IVy3ilwWHmsYwpf0nTgWc4yVXYritMsHaM2DZCtNhvIgvrgYK0dq1zq26uuzNeIgVaKzNvZkOFLUmkyfxZf+sVsdp1mQt/jlDP9cbjvTnt3jng/MuZzj6upJyMUc9owyGWFekx7+8zm24VcAAAD//zBMU69oAgAA")
	data["docstring/is-len.md"] = decompress("H4sIAAAAAAAA/3ySP4/UMBDFe3+Kd7fF3oojEv+5ClFQIFFeh1A060xiC2ccPOM9+PbIyQHSstC/Gb/fb7zDTWJ5hzGX+ckBxmqKh8AWuMACYyn5FAce1oSCCsPnKsYDlL9VFs/qPo6glB4jfKJUyRiWQThRin+isEACT4LCSy6GaIrEMlm4bc8JLETFWMVbzIKHmBIKWy2C3krlvsN9YIyxqEGyPD0v82tEazJEWRl+r9s2RZnQj5SU+8653Q7vBR++07wkdg6PQvY3z/AcL/DygM+v8Bpv8PbLwbn7Vo+38KV2zn2KXxlzVsNSeIiejPX2DKsJoKQZR4bwRA3g+KMNLCxD69dq76/28IEKeePSuKNiZhJtFm01nsdLR5rrKsfa+s3/XyfrVtKrFfU6cEr5Gnd3/wXcjOHInqoy1EqU6R8/onM/AwAA//9ltyBNWQIAAA==")
	data["docstring/is-list.md"] = decompress("H4sIAAAAAAAA/3yRT2sbMRTE7/oUY/tgm7oL/V+fSg89FHr0zZTlWfvWEtFKi95bO/n2Qdo4kBByFTNP85tZYRO86C/0KQ8ftlAWFVwdq+MMdYwxp4vvuKsKAWVGcYj524NCeHrmC4WJlKEJVAW74o5Q5wX9FK36FHH1ISCzTjmi1Txx2+DgGL3PoogpfizWm0ymoPCxxng+Mbt9PKPtKQi3jTGrFX5H/LmnYQxsDG5Q680nfMYXfN3i+A3f8QM//2+NOZRMPKtfRpov4sSWJuH6sbBNcaaHFxAubDXlxph//o4xJFGMmTtvSVl2r4AtRVCQhBMj8pmUO5weimHk2BWK8sd6sYZ1lMkq59KIFwxMUaCOtNac+rfWGKZam5bzdZWm0i9m/KXjENISx/3+fex5CfMYAAD//6e3brsNAgAA")
	data["docstring/is-mapped.md"] = decompress("H4sIAAAAAAAA/4yRv87UMBDEez/FfLni48QRCehoEAUFEuV1CEWbZBJbOLZlb+44oXt3lD8cAlFQ2tqdnfnNAS8mSYn9ewwxTy+PUBYtuFqqZYZaIuV4cT37daJAMrHtmE8DxPv9nxfxsyihEbJPQG+Jp0UlQK0rGObQqYsBV+c9MnXOAY3mmU2NsyUGl4sixPBqV9gHy+wVLqyGHiLbvgsjmkF8YVMbczjgQ8DH7zIlT2PwO+CPd0Emomqd99X98Uwcx1t1x5fXeIO3X4/GnBen3BT+NLpdQctO5sLVjFqXNzRwBYILO425Nuaz+0ZMsShSZu86UZbTXxQ6CRBfIloicBRlj/a2LCSGfgm2nHh+ekZnJUunzAsmVzBRQoFa0bWCOPyrqmleWeoiv1GoVyJPv5BUlt7H6r+ybyX9DAAA///YqAyuMQIAAA==")
	data["docstring/is-nil.md"] = decompress("H4sIAAAAAAAA/3yQP4/TQBDF+/0U7y5FLuKwxJ8mFaKgQKJMh5A1scfZEeNZa3ecwLdHuw5IIHTt7syb3+/t8GSiHzClPL86wLl4wS2yR87wyFhyusrIY5sooMww0fB5AqneH/lKupIzPNXP57po8CgF02qDSzLcRBWZfc2G3vPKfYdTZEySi8OSvTbR31NlVYdYA/iTsC2LXdBPpIX7LoTdDh8Nn37QvCiHgLvN/ukN3uId3h8qzyGEU2XhbexvlC0KZx5oLbxdbEhVDVJAUCnehfBFvjPmVBxL5lEGci7P/1gOZCAtCWeG8YWcR5x/1oWFbazs9cD+YY8hUqbBOdcapGBmsgKP5K3YNP2v/XltXXmLF+2a8UNTfoysmh7x9Xj89qLw1n34FQAA//8j90f0+QEAAA==")
	data["docstring/is-promise.md"] = decompress("H4sIAAAAAAAA/4xRvW7yQBDs7ykGKMD6+CxBHiBKkSJSSrooshZ7zZ1yvrNu15C8fXQ2WApKkXZ3Z3Z+Vtj0KXZO+BFtTN2/AsqigotltZygltGneHYNN+OFgNI4yiAxLy3I++uGz+QHUoZG0O1mmzkC1DpBO4RaXQy4OO+RWIcUUGkauCpxsIzWJVGEGP5f0bdLGbzChVHPzDIRuHBC1ZIXrkpjVis8BTx/Utd7NgbYNNyi381Gi2Ie7u+GcxT9Li/fdtjj4b0w5pDF88T5U/v0F0euaRAe5al1acoKTkA4c60xlca8ug9GF0XRJ25cTcqyvQumpgDyEnFkBD6RcoPjVwb0HJpsNb9YL9aoLSWqlVNOzgk6piBQSzo2EtvfuuuGMV7N9LcOyzGlxWx+adn7uPyT/ak68x0AAP//p+aZ8UcCAAA=")
	data["docstring/is-seq.md"] = decompress("H4sIAAAAAAAA/4yQzU4CQRCE7/MUBRyAiCT+y8l48GDikZsxm2a3l5k4O7NM94K+vZnlJ9Fw8NzVle+rESbCmyfUMTUXUyiLCnaW1XKCWkab4tZVXPUJASWG8KbjULKY1xrk/eHEW/IdKUMjCFvyrjpFZ7krQK0T1F0o1cWAnfMeibVLAYWmjos5lpZRuySKEMPl8f0Ylc4rXOjBTjX7BhfWKGrywsXcmNEIzwEvX9S0no3BwXI8ucI1bnA7xfsd7vGAx4+pMcuMxfvwOSpj3twno4miaBNXriRlmf3RKSmAvESsGIHXpFxh9Z0fWg5V5svY48EYpaVEpXLKvk7QMAWBWtJ+zlifW77p+lE01/8eV+a94aBXHFr2Pg6xWPxH7CcAAP//82gtTf8BAAA=")
	data["docstring/is-str.md"] = decompress("H4sIAAAAAAAA/3yRQU/rMBCE7/4V0/bQRu+9SA+49IQ4cEDi2BtC0TbZ1BaOHXk3Lfx7ZKdFRUKcPTOeb3aFjWi6Rx/T8KeCsqjgZFktJ6hljCkeXcddUQgoMUSTCwcxTz3I+/MDH8lPpAyNF8HfHBCg1gn6KbTqYsDJeY/EOqWARtPETY2dZfQuiSLE8G92X4QyeYULpctXyOzPqqYnL9zUxqxWeAh4fKdh9GwMzmDrzX/c4BZ3FZaWvY/LyphdrsSz9HujOQ57bmkSnn8tzTIlnIDgnWhtzLN7YwxRFGPizrWkXIivYVsKIC8Re0bgAyl32H9kw8ihy/3zB+vFGq2lRK1yyms4wcAUBGpJy8ax/+kYw1Qm0xx/3rwu5Itr9Aov2+3rr9TzHcxnAAAA//+lOibUDQIAAA==")
	data["docstring/is-vector.md"] = decompress("H4sIAAAAAAAA/5SRvW7jMBCEez7F2C5s43wC7v9cHa644oCU7oxAWFMrkwhFCtyVnbx9QMnOH9Kk5sxwvtkFVie2mvIftCl3n9ZQFhWcHavjDHWMPqeTb7gZFQLKjMkj5n8LCuHywCcKAylDE+gi2ZSECHVe0A7Rqk8RZx8CMuuQI2rNA9cVdo7R+iyKmOLnyXwVyhAUPo5lnkImv49H1C0F4boyZrHA34h/99T1gY3BM9xy9QVf8Q3f19j/wE/8wu/btTG70osn/etaUyYObGkQnr4e6xVWeAEheNHKmBt/x+iSKPrMjbekLJs3xJYiKEjCgRH5SMoNDg/F0HNsCkT5YDlbwjrKZJVzmcQLOqYoUEc6Lp3a907SDeNuWuIvp6lG/NmVf+44hDTHfrv9MLewTbF5AT6FVuYxAAD////9lvg9AgAA")
	data["docstring/lazy-seq.md"] = decompress("H4sIAAAAAAAA/0yPsU7EMBBEe3/FSGnsQ4eC6Ogo+AO6iGLtrBULx87ZG5S7r0f26QTTzr6Z2QE60u16rnyBz2U9nQy2kufdcQWh8mXn5BiykCBU8A/FnYRnRLqFeFVqGPCe8HHQukV+UwrQM/sEH2xPnb4UAOjIgskH2wpXOxMmgr17//U3R7uc2gbdIAv9BII1TQ+qOyNejDG91+dyZnILpgNa6JvxOo79qgU+OL2VkCQmHA37XNpX9/XoToUsDB9Klc77YHMi5wLSvlou9Vn9BgAA//8aeNH1NgEAAA==")
	data["docstring/len.md"] = decompress("H4sIAAAAAAAA/1SPwWrzMBCE73qKAR/+GILhb/MCPfTQe6HHIKvjSCDvOtKKNH36YjtQelx2vm93OhwyBZXXHoXWilRYJDLlYhE6bVPltVEC3UfK+ZHbFtLmkWWNMXOmWEUSeARtYn7Mv+iAt78upIrsv1O+I+i8NOPnEb7eJcSioq0eoQVqkeWWKpEk+GVT6oSRSS77lRWzmCqmJsGSCm7bk36FvIClaBmc6zq8CF6//LxkOoe9+L/DfzzhGae+d+591XBPPCx71fPpPLifAAAA///Twvx/KwEAAA==")
	data["docstring/let.md"] = decompress("H4sIAAAAAAAA/1yRP0/DMBDFd3+KJ3VIo6JI/GdASAwMLExIDFWHw7k2lhy75C5NwqdHcdJS8GL55N/ze88LLD0r1o+BasY2NvXTapP2VQ4Scbsg8NGSx4EaR5+exXw472EbJmUQAnfzDbFxzxcz5sIOWjFkz9ZtHZe/AtAIrUgnoMCrohslteIAPpBvR+W/8GhJ0DmtXDiDQaFEw9o2IQENS+sVcZtOnkSPgi6GwpjFAs8BLz3Ve8/GYI7fI1te4grXuMkN5jVgfYs73ONhs0nDpY3BkqLHkOfGvFdOwJPUFODUiXeisOQ9l8j6LLkkHNhqbE7zIfvv3tP3gOkRDsnylCQKQ/ir5WBZCrzFVA9pgrSLZ9WWESEquB8dxFbFlXysY/qYmoOmOgvzEwAA///hYwBj/gEAAA==")
	data["docstring/list.md"] = decompress("H4sIAAAAAAAA/0zMO2rFMBCF4V6rONjN9SVkDymyhtQiOkaCkWRG49fug1+Q9p8zX4+XpGYYq+b3e8Cv0hsbPApXHCf3k0Tu/i9jjbURFGYWa/BKWCS4eJm9MZxkw6R1SYHhA1WhtFnLtcuT7ZeURpR6zw/mefl0ru/xVfC9+TwJnQNegSM2dJEitRuesqOzSOVVTnXDPri/AAAA///Sf+9X3wAAAA==")
	data["docstring/map.md"] = decompress("H4sIAAAAAAAA/1SOzWrDMBCE7/sUA75YKQTqhtJrKX2D3kIIS7KqBWtJ1U9i9+mL1RDIdb7dma9DP3GErf6ELD9PBsq/ThdMHPOaVPEnyfSRhItk8MqXO8B1DFlwYa0rTIIyCpLkqgXBgmPUxfnvFscULu4s57ZWXPAooYHbe7CPZ/f1LVHX4d3jc+YpqhDhX7u3Hvv5gH6zwYzBGOyfMeAFu4Mh+hpdxtWpIkmpybf2R/1jP2CHV7yZ45b+AgAA//+tps20DAEAAA==")
	data["docstring/meta.md"] = decompress("H4sIAAAAAAAA/2yPsU7EMBBE+/2KUdIkp4iCMh0FBT0dQqe9eEMsOetTvOaCEP+OfE4HnTVvxs9u0a1ijDlua49NLG+aYIugxI4rQi3Qy3xH5QyfwKrR2MQNJda/M1vYavvmQ8BFDoM4cJmDU4qTZ/OfgmRbnixv8kDUtnhSPO+8XoMQAZ2TWeFivgTB2/6O7nTCjse+L/D+hQp7otfFpyqstuPNIcSb14//nSPRGfge7esqaOask/mozYBReZXj7gGjixOa5udMvwEAAP//4WKBkTkBAAA=")
	data["docstring/ns.md"] = decompress("H4sIAAAAAAAA/0yOsW6EMBBEe3/FSDRYCvxDihT5ghRRijUeEktmjeyN4P7+dBw6Ue7O6L3p0GuDykKPSvuv2iDH3VaZ6L5Szmdw/WMqatwNQRojiiLZEzPic4b98dKNhQ1aDJIrJd7APTV7QzJsD3wgpkoxRgTOpRKBSX9PLePoXNfhXfGxy7JmOgf0mYZvPda3snBQbsNL6X8cAPRrVaj37h4AAP//3EC5U+YAAAA=")
	data["docstring/nth.md"] = decompress("H4sIAAAAAAAA/3SSMY8UMQyF+/kVT3cFu9KxErDU6AqKFR3Q33kmnh1LGWeInb1bfj1KZvZooIz98vz5JffYqU8YU54hGvgVgUcq0b/skdmz8IUNhAvFwuivq6j7zl6yGnzireUTOQZS9IwxFQ0gb21beJBROGz+aYS4YUjqJMr5gNPYhGtbDKl4VdVaX53sdnq781CPChafOLfWRr3BiFX4kpUDUgYpOOeUW53EOBzwjXmBKGbRsMK3+XXZzPjNOb3vqSq77v4ej4qvrzQvkbsO2AUeEfFu9wEf8QnH/b4Wa44RR9y95KTnu33X/ZzEwOu9DcjwtPWf0PNAxW6LH7Gri3z2aS3sK23P19T4GKzhlsPfSKOYr4SnRh/wg38V1oHxuCxRBnJJ2nUnNWdqBsVEz83mWX16xlh0qKKHLYAA2ywMVoYJZG2MgTTgwoOnbO2lKVrCRBcG5XOZWaumTuWAIJkHj1d4qrPmw/9z+0dmLxIjrsJxXd1orvlZfWBaP92S+SKpvMV76P4EAAD///VqQ9/KAgAA")
	data["docstring/or.md"] = decompress("H4sIAAAAAAAA/0SOQQrCMBBF9z3FBxfRor2DC9eeIdRpE0gzZfJj9fYSK7gbPm8e74CjGia1pe9PWMXaWeBRghovY7SxRsY8w6m57vb0qXpKAYNgZyfTBUkmggqLc+CAa0FRzfAFmgXyf1N4OFpleDu0Wc7YYkowYbUMBs99H3BnENtiEUTu0Go6ijya5udsaa0ly4vfoKH7BAAA//9She7z1AAAAA==")
	data["docstring/promise.md"] = decompress("H4sIAAAAAAAA/1SQMW7DMAxFd53iI1mSIMgdMnTo2hswEh0TpUlBluz29oXgxkVHiR/vP/KIUy4+ycxn5OKpRZ5BhmaFZ9eFExbSxuGDays2g/CbhxjqyBi8TPABhKFZrOJ2w/sAylmFE1apo7famVSebWKrV9RR5j2OVVRxeajHz8sVK0kVe3YuaCtHdTwYiVUWLl1JCIRIqn3ULV5SO7SOVCEWtaVto1f7Dfc9Hcngpt//6W6RbyEcj7gb3r5oysohAKfEA/Lfwc79L+Mwsqofttc5/AQAAP//etNsWFIBAAA=")
	data["docstring/quote.md"] = decompress("H4sIAAAAAAAA/0zPQU7DMBCF4b1P8aQu2kgoUoELsGDBgh0H6DSZkJHssfFM2ub2KC4gtrbs9387HL6W7Iwp19Shsi9VDT4zrPAgk/DY7iCKkZyQ8sjhnUlFP+EzOaKYG0hH2JrOORquEiM0O84MvlBcyHns8TGLIdFQM8TA08SDy4Xjet+jxCBDqVxYx+17UvCtVDaTrLiKz9sRlWxec5kZh33Xh7Db4UXxeqNUIoeAX9PhiEc84bnrQmjbretubJNRnCvFJkAln7luJIXXtfEyqJSfPl3SmSuOD1ukTBDHlSuDMC06uGTtgTffaLZq1jXlxe7R2/N/kNP+r+zUh+8AAAD//+9CmPqDAQAA")
	data["docstring/range.md"] = decompress("H4sIAAAAAAAA/1yQQW6rQBBE93OKkrwBiY/AXxzA+vqLnCCLKII2NPYoMz1kprFwTh95EnvhZVWp66lrhyKSnBjeCjxtsDKWGCOTcgIhh+bfQzv6uiLx58oyMvRMiiVyYtEEPTNk9UeOCXMMPncWVka3JnvhEhoyouDtblU4Xm/IyJ5FaxycA8XT6nMhRUZY1AYhV2PwVgZMPNPqbrSAvukrDJ62J9vK3FcgmTAk5eUpbfsaLzOCuCuC8IMHm7DEcLETTxVs1pofn0Dph1Mbs9vhIPi/kV8cGwMAhYY/Fx41RBRKH4zuvmrbwMqMrixLY16tc4isa5S81e9J/9Y2aDvsG+w7/G3e+9p8BwAA//+GOQX8lwEAAA==")
	data["docstring/read.md"] = decompress("H4sIAAAAAAAA/xTKuw2AMBAD0J4pLNHALixxil2kSaL7CMZH6V7xTlwuIyL9xlbA7cXTY6FNCn3kRA19ywZF0NL2rpbliuMPAAD//w7hDINBAAAA")
	data["docstring/reduce.md"] = decompress("H4sIAAAAAAAA/1ySTY/TMBCG7/4Vr9QDrbZELFA+jiBx4L6cEFqZZNKM1rXDzDht/z2yk7SIU+TRTJ5nPjbYCnW5JfQ5tlD687DDHNHyyhRbUvfdSLyRIk0k8FAypP6esJ9rOB5hA7GAAp0omsJSSed4DAQhzcFK0uRDpgZPw8w1ThGjpIk76nDKajD/QrBzgpdjrr9a0lnU4GMHpTbVz+xwR1JsU45GQh28UDECRzb2YQYr/DgGpq7Y2eDtJlEZQr43kn0tnJ1Ls+U1Ck2csqL1oc3BV3FWZC0srTmL4aK9x3ngMEtEutiq+X/V0s1a1ji32eBLxLeLP42BnAO2HfW44Ocj3uId3uPwa7dGr3i1/YCP+ITPeHyzq/Flrw+44Lpz7mlgxZlDgJBliZVax4Hnw+G5ce5HDPxCSDaQgAu1mNQedR3AP8dS4mVGrEiRkGK41q0p2PQ28EC9LZhe0mnpdTmb284bfM0GrlMhr1dQTPk4lP0IjZIqtKJ+0+AnToJcjgo+3kjzkb1e5ztRa0ka9zcAAP//5r+52eICAAA=")
	data["docstring/repl-cls.md"] = decompress("H4sIAAAAAAAA/1JW0EjOKdZUSM5JTSwqVijJSFUoTi5KTc3jcgaJgAWSS4uKUvNKFJLz80pAdH4akjo9hZCMzGKFzGKFRIUg1wAf3fy8nEqFtNK85JLM/Dw9LkAAAAD//yVoAgJhAAAA")
	data["docstring/repl-doc.md"] = decompress("H4sIAAAAAAAA/1SOQQrCMBBF93OKD91Y0N5BsDsXIl5gSCZ0oJ0EJ4Xm9tKKC7fv8x+vwynmgJTfS4+oXmZujpjDuohVrpqNbj9cJ/mf9t9BvUjQpBIP0xmawNYgm3r1Aa9JHepgPMfH/ZJtbkirhd0xEHUdroZx46XMQoRvE3uz0NMnAAD//xQLeSyiAAAA")
	data["docstring/repl-help.md"] = decompress("H4sIAAAAAAAA/2SOS27jMBBE9zpFLWVh4APMbuDxbhaTIAcwJZUkAs0mwyb9uX0gOTFipLf1XnV17UJJuw73++stibuhLN6wJgg0czObrh3jgCnmsMFf4BiHGqjFFR91j+PVhST8jdOGz3F3arq2GqH2+eSwOJ2JoeZMLVAXaMkNfJJXY4p3exB77MNB6DLKQtiQSW269r368gBeqi9b/Hr8/69p3hYi5RhSQU+JF/iRWvzkaRv1Y8UvTFEkXjiiv22I1tAzI07gNWWa+air7AouXgSLOxM9qeDZSXWFI7w+ta9bYHd1jz8wr7PwWx8Gp7DkFKFK8UkI8UrbNx8BAAD//3pr7pSgAQAA")
	data["docstring/repl-quit.md"] = decompress("H4sIAAAAAAAA/1JW0CgszSzRVACRxQolGakKQa4BPlyeubmpKZmJJak5lQqpFTCp4ILSkpLUIrASPYWQjMxihcxihUQwXzc/L6dSIa00L7kkMz9PjwsQAAD//+kUzLVaAAAA")
	data["docstring/repl-use.md"] = decompress("H4sIAAAAAAAA/0TNQQrCQAxG4X1O8cNs7MLeQaQ7FyJeIITUGWiT0mSg3t6FgtvHg6/g1ENhMUAq20sDxqvGxqJ0/ZWsCun7rpYQt9Qj/9eIZ22BFmA8pvvt7La8MXeTbG4jUSm4GKaD121RInzF2X2gTwAAAP//xZr5jn4AAAA=")
	data["docstring/rest.md"] = decompress("H4sIAAAAAAAA/1SQvU7DMBSFdz/FkTqQLBmgQnRCDAzs7LXjHBOLW7vxD83jo7ooFZvlc7+j+90dusRckLn0SCw1hYwyE+03uvbOXCqDpfqcfYarwRYfAy5e5I+B2YZQZlPA1UqdeOtyPuUCCk8M99IzrXee00YO+Ci4xCoTRmJkoPPWG0GJsDPtN1xMMPgx4u8UavbhC7rLXF6bh8ZIFxNhjUjL2gIaMUFfvfSg1G6Ht4D31ZzOQqWAbqLDiofucMDzHk+P2L/0/TVop1h7pdD8eWP+6R+7DTkO6jcAAP//LTZxf1UBAAA=")
	data["docstring/str.md"] = decompress("H4sIAAAAAAAA/5yPS2rzMBSF51rFSTL5Y34CfWwglG6g6ayURLaPa4Esmasrp9l9ieyGQmcd6ug87rfBv6SCLspQVVs0MUwUTUVI0AiLpOLCh3kSWmWCReB5EdFJHKA9l7frHFtM1mcmxK78jBIn17KdK3fGbDbYBzx/2mH0NAbzBeue3sc13u5wjwc8vm+Nee1dAmcjzs57CDVL+LGIY8ndUsd54IW2peBQPMmYvVdKsOom+st/xEA0NiAnlq5TUlmd0OXQqIvhyj1KbHPDG/+CRknF0FtFNdhLhZqwtec1UxNC26K+lNbDmFUpRaPsvlFXf2f9FTzuzFcAAAD//9gkDpPCAQAA")
	data["docstring/take.md"] = decompress("H4sIAAAAAAAA/1SQvU4DMRCEez/FSClyJ1AkfhJoKSjokSgQipy7uZzFxQ7rNUl4emRffqCdnZ1vdyao1H4STUheEfl1VSMLEdoTnZOo4MANvUaELjsSfcNo3twwQKhJPCwG+3M4D7ORTnsKpiV4iiDouKNcwjoJmwJZce28d36d17KwlfDtWrYXGKom+MYqvVW29QwvozOGJA3/+KwQsQ+iFGhvfXFJHkdlO355fRRjGjRTz1fv8kcrQiUVWDszZjLBk8fz3m62A40BqpYd9phWN7jFHe7r+iQe8D7HAg94/Cha6XWBPQ61Ma+9i+AYM4KO1eVb/pe3PEVjjkW9nJnfAAAA//9Y1ktVpAEAAA==")
	data["docstring/to-assoc.md"] = decompress("H4sIAAAAAAAA/2SQzWrrQAyF9/MUB2dxY24aCt1510UfoXRRClFn5Fh0ftwZTWzn6YtdGgpdCaHzfRLaYa/pjkpJFoU//7ewKV44a1nbytFygSZQxBYSUrkwiuZqtWY2L+L9ylhSjqQMQmFF6n/xEv8YjniOXj4YOqw2io6yw+lbdEJfo1VJ8YBpEDtACjxdxS+wKYxV2R02MnOpftumg5Qbhmm96p0RSDkLebmyg4TATkjZL0djdjs8RjzNFEbPxgB7xz1mvHaRAqMpY1Xl3KCjM+P++PDW/oQW/Nt3E8t5UDR+LU27DW+vnLG05isAAP//+ZF5XlwBAAA=")
	data["docstring/to-list.md"] = decompress("H4sIAAAAAAAA/0yQTU7DQAyF93OKJ2XBjFoqAb0AC46AWFSVamYcxWJ+QsahSU+PEiBi6efvsy03sFruo1RF5c+dgy/5iwetSzly9lyhBYQFMW8S40J4Us6kDEJlRWn/0ZI3/oDXHOWDoR2jKuVAQ8DlZ8AF7Zi9Ssl7XDvxHaQi0k3iDF9SPyqH/WoOXMe4btFO6qbhulzzzkikPAhFuXGApMRBSDnOB2OaBs8ZLxOlPrIxgA3cYsLpAY94wvHs/rLZAIBN1MO2GafpDLvDhKNza+fO/ipuDbavTZid+Q4AAP//kawXNEYBAAA=")
	data["docstring/to-vector.md"] = decompress("H4sIAAAAAAAA/0yQwU7DMAyG73mKX9qBdIxJDF6AA4+AOCCkmcRVLRKnNO7W7elROwl6tP/vsy1v4K08nDhYGVD5575BKHriwepcjqyBK6yAcIPcu6Q0M4GMlYxBqGwo7YoXXRl7vGmSb4Z1jGqkkYaI423EEe2owaToDudOQgepSHSVdEEouR+N424xB65jWvZYJ/VPw3m+54uRyXgQSnLlCMmZo5Bxuuyd22zwonidKPeJnQN85BaTAwCfqYdvFR/TJ/x2iwmHplmiO/+IA57w3CyN1aP+E0yN+w0AAP//fvP0f0QBAAA=")
	data["docstring/vector.md"] = decompress("H4sIAAAAAAAA/1SOQYrrMBBE9zpFkWziwM8d/mJuMDDrHqmEBHLLqDuOffsBx7OYbb3i8a64rYzeB3If8/0+IQ6K0yBQvvCG4au2dpI/AK/SjWDjTHWDDMILwVXaU5zp0BqW0deamB74LNWQnxq9dkU1aEeqOXNQHV5ED8Gpb9U5pMF2ddnALXI5Xo7qiKL4JvzoSqgKQa7D/F9sYoYsVmrXRwjXK/4rPjaZl8YQgFtixoZLYWv9Mv0uOy5eOPhezogN+xR+AgAA//+vfXdiKQEAAA==")
	data["docstring/when.md"] = decompress("H4sIAAAAAAAA/1SQTU7DMBCF9z7FU7sgqUolKH8rJBYsOAHbOvGkGTEdV7ZDk9sjO6WUlW1974013xLVqSfFMZBD58NhtarRenWc2KsVmUDfVgabKMKiEd9+mY8OqacLcKXNrU0EjkhhSP2ESn1CZyXSGvmqLPW69PI3ETZcTdjgk0Uub5Bt+5IDK9IQdI1A+WTdzzNY7SXPXmEjOEUEioOkjTHLJd4U76M9HIWMASpHHUbcVHe4xxYPeMQTnvFS14UWC9UrKiHFWGNbGwCojoE1iWIx5t0a3i/OYM7V5tdGnA6NF+zyoFv1aZfzQyQH1pjIOvhupruiQUvrn7k/oVZdwcU3TtlNcy3cq0zg7rxu7mbFs3nzEwAA///KLwLw1AEAAA==")
	data["docstring/with-meta.md"] = decompress("H4sIAAAAAAAA/1yQsU4DMRBEe3/FKGlyUUhBSUdBQU+HKDb2Ol7pbn3crpNIiH9Hd4kIoh17tO/NGpuzeHkY2Am5TgMGGkdOHcidYmHD/JTICV5x/RNeM7wwbOQoWThdm2KINNKhZ9SMA4seQarVyTnt4EUMuWl0qYqz9D0m9jYpCMpniJqTxqXshRwn6htjxvt37Zfoxpj2eCt8j8UwTvUkiRNEl/ICWDPo5gfjz8YaeQdrsYAMpCCzGoVcTgzzqUVvE+9DWK/xrHi50DD2HAL+jBYAYJMV75cPbLZbXPDYdUv69ZRqBFZ0t168Um2Hnm0eRI+2+u7CTwAAAP//PWKwV4YBAAA=")
	data["docstring/with-ns.md"] = decompress("H4sIAAAAAAAA/3SQsW7yQBCEez/FyDSg/wcpKelSUKTPCyy+MZx03rNu18G8fcQp2EmR9ttvR6PZYHuLft2rQWUg+lyGfzvwU9IkTqvAIBeJag6plo3SsTktjsDoyP23HBV+JbqszrlygY3sYh8Z1oAD3vsqLgQh06DZIalQwh2co/l/RMctpoQz0RWKM+DMPhc+i0a91Kha4ICP63JhQKFNqfZ4KEnMq7dEFvpUlOHQNJsN3hSnWYYxsWmAbWCPGW2e3GLg777t7mE89xvu+3UcYP2N+scrsB2Loh1LVMfLES3mXc38gV+PD9p8BQAA//9UbL2/qQEAAA==")
}
