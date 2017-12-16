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

var data = make(map[string][]byte, 87)

func init() {
	data["core/00_builtins.lisp"] = decompress("H4sIAAAAAAAA/5RX23LjNgx991dwPdNp046m09e42X+BScjihiJlEspG/foORNmxCGqszVuAc3A5ICn4dDqdVBpGIoxKh4iv6jxaR9anw+EPg22z/Ktsq5RSryboBlJCUkfbHpV6TQNqC65pQ+wVxRFf1jyHVPAc0vE5z4Qynwl78kUEs+ax5Vig8APcGsWWEmVTg9d1z2w5vhwOp5M6Q7JaDRGN1UAoBEsNUOhlSG+dNL7j9DNEkyP3SGCAoIj401LXsG9V0t1aVj8jValiDWnTHGBJDjqGshmD7WwX4W6OHRO9joFQKVnS7Diq15yBmc+DpckTfDY55gOz1IAd+DmAN/9su7Y9DbjKsGZAVitN/Tm4Uq409WWTaepL2S/oGfiIyqbKfHIeaXdBg8uleOgxDaDFSZxPiE/y1Pi0Y2w+yYkxMRufs5thpJLNto3LPItK0fqLEJWiEJViqVWi+O05ihWluGTD64heypbwKvLNl3+Fam1MRX+zqcRFTOVDyKYS5kDA2FTCdJjH8ghjUwX2Q8J+iKToy04d+hLlqRPngDp5qKlEXVDUDykFvUbNJlGZlXJYKccHagpxBcsmfqrF4PFauUfopdF6g59YPg/8sHOttXdhGGpwLllac4XL3R17jFaXxXo996SO1uuIkDApYOwZozpPKngslTB4oxjcSflrEVgdwRhGs36hXUjiTDV3eBrPFEHTc86fd04/OrKDs/ic9PedZOyHNTsYv90ZQwxm1JgUdagi9sCjjKoNUYHK4YwMto72dgumO9TvMvscDa8jOEtTWcq3tx1k67fo34vc1AHJGF+ngu8AuFHM9vvbnjgmYPK/k7odmc1w/+4q6+vkbcb5pbJubT6GE7dpCHxf28q6hZcNB/h8+drRa7JBrLzahTRGlJ+/xfGLK4uD/mygsv9kx57lFobBTbUVanaUKg8QyfKeW6IXh1gY4R2bW9PbOxW/f5yu8kl9KD9Lq4PXY4zo9XSoJLuEXUuE7sBXmmaz6DmG3iYUHWdzZQ1YPLlcB/9Nm/vAXDEj+AvyvOg2xAZBdzsXJVYKqDbb7JF7h+PfbBV89sjhDurrb/WLAAa5rPDrWQNnT4kneMd6cPaIb1QMQx3NnuqpjOAvy4wi8HzBGxVRhw+MtTFhjPxpXTcF85yLRucQm5toPrykyPZYvg7Z1bDr5fB/AAAA//92+qEYSQ8AAA==")
	data["core/01_basics.lisp"] = decompress("H4sIAAAAAAAA/3TMsQ7CIBSF4f0+xbGDKUNfoCy+hzER4TZpYoFywcXIsxuxgw6dz/k/rbWGxJIzJ9iQeMTNyGyFqHc8LcamAMcTAc/RBTsYEc7oHE/di4CzNwvjYe6FLwRcey9DLBn9Zo5eFGo71fZS6lfm9WO0AUfYsET5MrMMvG4J6qkt/+lht/UhYwdQ9A4AAP//DfilB/EAAAA=")
	data["core/02_branching.lisp"] = decompress("H4sIAAAAAAAA/6RU4YrbMAz+n6fQOjhsWOB2P68M7j2OQN3EaQ2OXWyn4xjLsw/JTuo66Ta4Ql2qT5K/T5a03+/34C9jCNJBa518haMTpj0rc6oq1sl+EK2zYGyoAH69drathfcywM7YsPtdAbxfhW4qgANTPUxXoaEX2ksIbpSc50lGo6X3ZR7VUxr2HqQPDdw+Rmm+ABDO0iQ03kRWozRMCPHCFaT2ssldJ7TMzne8fp6lKVmhbZPXPaveuqEpOE1oJL/M8Ylc/R0j1lmY3sjOo/+KVr1R+E9Qo3Ih8C9q6JjTK6i11nQlLbQlWk1JqNVixOeIv9H2lP55bB4ApmWASIglAJ55dWuH2Jc38DtvEspiYwETpgOmfC2CHQBT8XguWfJXivmyCxgqgOmNOSJBZ7qM9N8TKUsi1hURq4LQTPxvRQ6xJMJ0X2HKarLIVj0QSLpn3gtjwkqWX/5C8/7+AzM2zKlvZciTWVemsu4THZD0WvdYLmL4ZdatBa+0PuS3qZRSbgtVfa1lIM5HZTplTmkdYT8hBNNsp/2SDX/uH3dS7HYk5EIt3Mkv+nx9lW2wDlIQh90cnoBh9AGOEvx4uWglu90c+wNecILMEvogtrUmCGXgBaSWgzTB73g2fv5jADbHPfNv2XTEyZyx2/ClV5swlCarWYJo2KId1dMCXr0SbbmyuvPu2qhvtr1m2xMcbfex9RpxhSFK1/4JAAD//y/68t3uBgAA")
	data["core/03_functions.lisp"] = decompress("H4sIAAAAAAAA/3ySz47bIBDG7zzFpxxSOKTaJDdbqvIekaVlYZxaNRAxuJeqfvYK421YxQoHCw3z++bP57ZtW/B9SokiTIjUoJ+8SUPwLIS01DttYoBmppgOOt5YAPLaoRw/jGoJmFFPTB3kXfvBoHHErG+EXUWiJDHcxAkfhLseItmdKhL7z/dOZOl3aYK3otSZ5Sf7piBrzfkiI3FC+a5ZapHMp6GRaW32qbmH6lHlU0/cewH8aWwwh1wuYdf73V8BXPfoQ3RLl3IcOOHbusDG6V90MGPgKdJSvx5Bam8hf0CO5IuCWoYZ+PCbTAoRconiTSn1uL+gj1v0saKPr+jTFn2q6NMr+rxFnyv6/NWBa6f+b4QfGxu1+7B6Ff1qgKVnC3KsmOC1I1ROvGcQ8xKWvV9v82VL+XsWSDreKMFR+hks9sg/UxGSOT6v73NJUJgvOUMp8S8AAP//CMSxMzMDAAA=")
	data["core/04_exceptions.lisp"] = decompress("H4sIAAAAAAAA/5xW7W7bOBD876eYc4ErCZyB5n46OCDvYRgoI61SXWhKWdJpdUX97Ad+SaKsJEb9Iwi4y93Zmd2l7u/v72H7s3PEqDqmPehHRb1rO2M3G1FTc1IVdyDmjjfA4U9UWp0t2eMG+CpO6pl2wQihrO0qXB6Sg5RyHqFXpq1WIrBqLUHEGNeXDVq7q5TWG+Dnvuf2VTmC4zP98sHscErhfDShTA3R2p3uKqVhh5PcIP7CaWtd8p7O6cX7QTQtT9Zleld92z22pm7N0zqOpuNTgeCVKtcx/PmU6x/8DaHJxOMSW0AsvAF3a/nX80a86BVbqgsAnjN8Dleval7WBGFc9vLZR0fTueBt6QVi/6i76jnlkguMTWuU1sPvoUyXr3CW6bPXCOANv1jzR17vFUM/er61kj9869/C90qhU1LHw65n6snU64ljRjzT8L3jet7zVWf+TYBwyHZ/arOi4okcygBSHhfJvfm9msO4irPRZO1IY57WUOnPxKmQSBr4/3LJQv4KXkKTw6EpB87KvzJZABiC6Q1bDzHCBctjMvl669FtbHA06KU/mPGLPsNrZHEjIw13ljdGY3nHd0pOcnUn0jG7sSdtCbOfCFsR+xNZq54I25PSfgVQHUSJQ5pSb4suncyPrMz768H69b2zQ9pQ1hK7neInu8llzLXENrIT48JPTa9apnrrC4kbX6v/Bn9ppmjutbdlfVUc3MW4aHBXOAQykcSft4AsdLaTAmGjf04P2F6T26D4HXzOWfGFrVFejTgpY4y6CzDkcabzxHUusmjPnOANfaLb7fqsEPz9G5n1mbudfaBnqhf8f5ETKyW1mdxwKRc4ty4bEIta5DodH6+YUA8x+7Vl/NO8JeZtAvpVaHV6rBUOF2Jen/7Lg1B9rweEAq7VmwGVy9k8eEgIsZfwb3jhbBKwT3ORoKZlndUpOzsiGD975p9MjoerD6Ybxvdqc4Dp5dwyWSgHTco6dCbzEKc6kJ6eiPmGve6kuNSwfEEnh5Ca7PIhnhzGF2H5oo8SezBM1SdPVdW9Eo/DWne4PITE8ojLpG3OKsueJ+YY5BO+lBYmmy134xBcHsRM6Qx0nLS2yQHDN6sPIeNfL+P/AQAA//+nk61XTAsAAA==")
	data["core/05_concurrency.lisp"] = decompress("H4sIAAAAAAAA/3RPTWrzMBDd6xSPGL5vBM0F5E3vYQJ15YlranvSkbQIpT17saxiJ6HaaHhv5v3UdV0jXFKMrPCi7OBl9kmVZ381hjo+T61XQS8G+HSd+GMbAkccejl8GaD5h1fpricD0DiEiP9Fz03tOx/9KCEpo1kWAPIyh7uVXrKCtfbGkGfWNvKDbcHvzV9o5IjGv7VzBVo++5Qt81ticAVyeUDe2vM8DRHk8reSJW8vWIWVQxpjBeoE38858Gm7B60OdgeVC7ti5AJ/FO3bpucUkz70XNE/Wl5UpiEXKtM+7kZuUUuIX8pa8xMAAP//D2AMr/sBAAA=")
	data["core/06_lazy.lisp"] = decompress("H4sIAAAAAAAA/4STza7aMBCF93mKo7tonUWkdgub+x4REq4zuVhN7GCbAq2aZ688zg9YpDc75hwfj78Z9vv9Hn64hEAOyjraoZO/7/B0vpBR5ItCNNT2UjmLYCvpvVUF8GfXWBV/UcDbXH/7WwD1l3jYHwrgKOQwdHewCKGsUTJgfI96WZZZdKd9eJEcy1vBUfs09xepYN2L5CRsZSf1/+lOmg/Kk7nIqaI+4PE7il7+pIoN+AZtWnwvSzb28nbYNI69vK1ObbC6n5xj1F6Z4QMNhy3zGNXsZXEJKk/n/HFzfab2wzb3SE3wKL5Om7Tje1Rn/cUR6mgAk/SZZY7jnKUFgxDF60l3kW49OGqgbNelm9bWAHE9kak6Cqj9YlkECD4qWu18gGcok8zNLELxNCeI9X6kBEcpgL9HUC2vVu3pXNFtcB4rk6P4IENOBkqtttZVJNUJ4+oW1OsA0ViM7xODlcIgXdBBW5NPYRGmNVP2YkJ6PsQiYi7P4rQUqRCH/kDsGevCT3seD5/O4U1/zkQrv+YJZ94S3y0aZ4e1jZntvwAAAP//r/v8qpIEAAA=")
	data["core/07_io.lisp"] = decompress("H4sIAAAAAAAA/9RS3Y6bPBS85ynm209d2VFp071MtO2+B0KqA4fEqrGJbYhQ1Tx7ZZsQ6K/au3KB4jAzZ8549vv9Hq7rvSeLyljaQb41WcZqajQ6m7eiyy/Sn3ItVQZ83nVWDsITvO3pSwYUTa8rODqXGcBa0WUAwBqNYhCqBJMNmHSBj0EoHl5gkRSOnEe8ozPn97EoHtEY27ooqsijcHQG+8YQnLf/JSAv09xpWkAHzfgnwN7cltxtnK9N7zfYXaz0BNZI63wywCfvxuYkqhMKUtSCWZoA5e/lNq4TFW0QmJwvd9LrpUTXqTGsmuxnv1bVdFFS02YpKLX/g5z+lZik9uonUYWF/zYtIS3V+UCVN/bDj4sc6hrH6TqGk8CptOwZ7FVIWKfzE8d2tt2KyhrEsE1HGsVB6lrqo8MjDqYek6pzZH0u7NGl+NaWcONwPNyVZqG2dx4HgsAnGvNBqJ6QmA8xjMroOsk+J5d3ue3tdj+y2uD6EhxNN/j+O/DTHRzbdGWzhS3H4vSOv56Q01Mp4+h/sCN5rFm7+AlaqrkZAPN2XAiwRXrXl1Sl9J698XJlfuI1UgulxpUXdjmRjleYmpOscbDpR3qyrwEAAP//cTSIgPwEAAA=")
	data["core/08_os.lisp"] = decompress("H4sIAAAAAAAA/1SNQQqDMBBF93OKj0KZQKXtttn0HiJUdATBJDIZtz17MVKwfzOLeY/nvffI62YmiiGpPJEyEY8yhX7QBJuDENBeMCUNuSPgzYsY2my9Wg3wsKlKtGZH3ZXwm0reFqvBY8LnVfTzW+JY7/ff7wrAq87Rlgi+gZsDPXoOj3uZQxVy5ehUco6+AQAA//+xXZhQ0AAAAA==")
	data["core/09_predicates.lisp"] = decompress("H4sIAAAAAAAA/7xWX4/rphN9z6c4v/mpd/GDpbaPu6r6QaJIZe0hi64NDuDcTavms1eA/+0NXrVS1RdjM+ccZg4DycvLywv8MIbADo11/IzBcasbGdgfDqJlZaB9zVc2OF5lN/IJ4heIH5A+8HOFH6tqBdq23cP9VFWz4rJG3cnfb7XnywH443lw+ioDI7iR/zwARzWaBo3tutMBEBssILSC0D5+J0SVZiOKA44KInGF0s6HDKhOEyST1fIFiMYaDwXxmBiyjuNFptoQn4SSned18knE5KtqrrWXjbNoWdWr9GD9J+Ua2TN6DrKVQea6U0WjaeoUE/7WQ/jgMpR+pbm03+KKuC/Q2ZJvOrzVUXI1QJkNDsf7UzbqC+5PsdTT1pxO+lC05p69SebNCplfbW26z9V8bovh879jC/2P/jNrYg5G/T+xju8nCGMDRDbmvdr23OdWJo0dI/+pkx89+/LRNek9u1BLd/bTOZpOt+jYYNUHze/oRx/wyhikdtxStZjftxDB1tJ726zUxXA7m/zQ/lPn3FOC976tykjD5wfkUviKi/eODLYHxSdNXj23tompcQBNAKoKTKM7kNEdLbv0HTPGSsTUaj64XWKMFYl8AXm+7BP5UiTGHaKOzS4xxkpEbVp+5xY0vdD3xHm+RM7bS2mgorUpUqL2chjisnmkR+oUKBarfQDF586OplCJeeUmWAfKY2HVKVBspMHZXnsGTS8PVs3zxXo5yHxwSjm/yYzYWdjGbVBxjxRlIqh54+arh4QZ+1d2UNZhsF4HfWVoo7TR4VZuaz5nuTrp7csZPsu/IScNyMi59T7IhTcZIKcfeu0Rr8B5iXKt+dKiNMy5BfbBQ6utkkSGFNtyGLobKA27Go008eaKIL3TZ37gRsuuVtb1oO3XR9XolcQUR4oXa/vKt2/WtaDp5fP6ZlAxs1v/ajtQHmmxvaw0oYpHyTayA6WBNvtXFDIYzWWUnVaa21W1IJt+Nig+CzVOHaY9EqCUVvyvSLZti5Wt/IioDn8FAAD//6EmouqqCgAA")
	data["core/10_threading.lisp"] = decompress("H4sIAAAAAAAA/9RU3arbPBC891MMOfAhwTF87WUCh7xHMFSx146oLBlJThtK/exFPz75cdpA24s0N5F317OemdVuNpsN3DB6Txa1sbSGP1gSjdRdUbCGWp0DpTelks4XwLf1YOVReIK3I30vgJ0XtiNfFQAbtSLnwKSL9UgpXgAAW0TymadmvaitQfkWmzSmLoVz5LHKn9BK6/wqNGS7o1AjVYh//BzBf2iN7V2V+5HHToFdcwCLSKmS89dYm37tnFNXYQtmKUarHP3EyjewqcWU+k5byzFtU1kG5re8fkZMiafilYltbSb3iJgzPS1Ni9HnsixiH4V6ycSq93r25UC61MbHwdVShc984Rd4LLGM0iSI+47fk+a+Nk/l+l+Q5jw2seChNsKdh2a+Cg6jk7qDwN6MuoEWPQUADMZJL40WCq2ir3IvlfSnS/li7R0NI8RvCHmjzRRxFuIEEpgUcnpJeV6jvfhMZW10U9ZKjI4WixRrN1AthSrD2+fd6k59XKytxi69eslisNSAaX9AyuF//hrbXwU/XLBhed+vA+BRqApMtpgS0OROfZAQU8Dg4bg0LrK4sC48Z3PUCbOR/mDN2B0gEEbetEkUiJAhaZHEFrYbe9KPL0JiMjsY7pH1pbCdy8TCcNKRdNBFz9WcY5WP6EfnsScMQlpqVmmA4xp/X+GsFwPYrVMo3zjYIKyPJPHxjH5XmD9WJiyGf0SYXyrzIwAA//97X1q6XggAAA==")
	data["docstring/and.md"] = decompress("H4sIAAAAAAAA/3SPMW/CMBCFd/+KJ1iApkFtt24MzB07Rqfkgi05dnS+EPrvK8dUoQOZ4tO79323xY5Chz7KcDjsMbLk3wRCslH0tXXSTk5duOBIoTua85X8RMoJahkl3Esc4LlXaIS4i9Uap4QUYwAlxMDgdS2C0JNP/IM85Aqz8x7COkmAWtIyr/GllmV2ieG0hEaJLXOXS+6N2SybBL7polMbs93iFHC+0TB6Ngblxt0L3vCOj71B+RaJv8cmRL07cLfZG/P9INUs0aZC4CvLI7r5v9dUmC0LU/p8zlWZVuwzZL4pqRTG2l+b3wAAAP//MxDI/rIBAAA=")
	data["docstring/apply.md"] = decompress("H4sIAAAAAAAA/1TPTUrAMBAF4P2c4kEXJggFFdy76A3cl7GZ2kCaxPzUentJgwWXM3zvDTNAcYzuB2v1C7J8abTZSganz7qLLxklgC9QbPA0HewqF8komyCmcFgjpmWr+EXA3twdtmQ0LlfJP//XBxVS9wt/ONHgHrvPj0TDgDeP6eQ9OiEClJEVJx7UE57xonVb9T8ecWqi981mSPf4ts4hSanJY36dR/oNAAD//8CDF8f1AAAA")
	data["docstring/assoc.md"] = decompress("H4sIAAAAAAAA/2zQwWrjQAwG4Ps8xY9z2YRNWNjbshR6KH2BQs/KWBOLjDVGktP47UudQFvoUeLj/5E2+EXuLeP/mRdcqM78sNttkY0p2EFQfsNKhEIuDA+bc8zG6VVqvcMf3JGcezRFDIzJ2kV67nHmZb+2YCIx/41mMI7ZFKTgcYrlW4oUaENpNjrIPoMOeBnEUWbNIU0h/uF6KYWNNRAD3Zq/hlUJNqrwRYOu4GvmaaUBCWRSHBmxHtRDFIQi5rHPldxRyAdpekhps8Gj4ulK41Q5Jdx/mADgn9LI6HyaI9i6245ODPw5/L1NlfSE7rl12/QeAAD//+eMAcV/AQAA")
	data["docstring/chan.md"] = decompress("H4sIAAAAAAAA/2STwW7bPBCE73qKQXKRjMQP4FvwI6f/VvRWFNCaWllEKdJZLu24T18sZTtW65OhHe7OfEs+o3UTxQ5OmJQzCCX6wQs79SlSgJUjh+bt9g/eVAMpIasUp0UYOpFaoWQeoAkHjiykDEKg3xdk/igcHSONOFEonLf4PjGEcwm6tJwoT68zHeFSzD6rjweTU0TPs9ceY4nV1QsIvQsp8+pbHED3QVt8YxXPJ+tCERx45qgYJc3Qib8czXTBZh+S+7V5wZl8nTsmqarIn7oYtlR7hjlRHpDkrrlxWQTV17DFuwnr8K/zdNeefQigkBPqZJSoPtRu+zKOLEZkDCVPPIAMzhVUGuHVyFY4J5ZVlm3TPD/jv+uM//mSm2azy/yxgf1WsRe3tz0NzWZn0arQcNWYcsdrg+34mGRG31p5idX1zWZXM9tR+jqgaUGxQtS39VvXV6NvEe+fNB8DNw3QBlb8cNP1Rv5szEp7SGi3cBOqPTyNKT11tVTLj6U9yb+lZV7XLN00vZ7YaRK0hgVu6qx2JZYhfBTOdk/MdUjn18AnDihxYDn6GH085BuL/gbv8RaeJ+8m256PmFNWOMqcq/4oPLIIDzjTZfVIVk/k+jJ4IQPapxPDpRIGW5jwWWw1cVeZPQS6dVuS/k2rfWDUdc2fAAAA//8pTMTx+AMAAA==")
	data["docstring/concat.md"] = decompress("H4sIAAAAAAAA/0yOzUrEQBCE7/0UBTmYIAr+3kV8A28iy5AtzUBvzzrdcV2fXjKB4LW+r6q7Qz8WG1PA+XU5QNNv1jPWjJaCvpCZNtLlubIlafHOG8BpKs6lFLRAdsREVPqsgfLxby3bZ2NUHmjhC2UaJxxr+c577rfNa5Guw5Ph5ScdjkoRoI9ypdlj+/ntBre4e8dFf48HPA7DIPI6Zccpq6Iy5mrtXmvt+qZjdXfyFwAA///8jynh/AAAAA==")
	data["docstring/cond.md"] = decompress("H4sIAAAAAAAA/7SSMW/bMBCFd/6KB3uxAsN2hgwJghQdWqBzuytn6mwSoUmBd6rkf1+clLhFgY71dsZ7j/fd0xobX3KH575yBw2cX+7uwEn4U4Oe66nUi8AkUWPJlHCslH2I+ey+lgomH7D/w72HTzQIb22E/R89KWOMKeHI4J+UBlLutqDcIZ4QFVGgddBwxSYXRXuiJNxuMQ85praZ0/b2wh62k1luUXNSZR1qttyigesYhWdT5knfdzJTX4tnEe52zq3X+JzxZaJLn9g5YNPxCRMeH5t5MmwHAJtnTHg4NMBqJEFiEWigjIfDahG8YML94dAsgnNlUq6L5v7wLvr9m0Ux48g6MudV49y3DA1R4Mlu1/6taG8H/ODc4XuK56DpikupDGLRwBo9pXRFn5gk5jPGMqTOfFpgJ6CMJ6sXb3wdS+2e/gfp8sI/SH8EgyApRky6YI+lvon1c2Q/V2XVvc45rx+r3goXg1k+mJ37FQAA//9CQizNwgIAAA==")
	data["docstring/conj.md"] = decompress("H4sIAAAAAAAA/2SQvW7jMBCEez7FAG7OuIOAy3/rIm8QIEVgWCtyZTGmdhVyZcV5+kCR7SYt55vhzqzwx6u8o/AHWs393zUohAJO3LNYgSloVkcWz27zW5vtGoWaxFeuwksXCxru6Bg1Y4opIcS25YzAA0uIsocKrOM5wGc2hp0GrrBBisUWy5B/4H8gHNnbJYmG5XXqYmKQgEpRH8nikdHTgQtEsR8pkxhzATU6GjQHzlH25+vaUbxFlSVT1DBpPmCK1iHR1+napqCMvgMVqHDBkDWMngOaE+qehhqaUbcxGee6cm61wkbw/En9kNg5nBd++48b3OJui3s84BFPa+de558z25iXKc4ldxf2gm53lfsOAAD//0Od9nWqAQAA")
	data["docstring/cons.md"] = decompress("H4sIAAAAAAAA/3ySvY7cMAyEez/FAFvEFg5r3OWnT5EuZYCUB57EtQnYlEPK+/P2gezsFkFypYYc8iOpA9qY1XHKNsP5V4eY5zdRdpCCJ55ZCy5SRlCNr6yRm5/bW5EtsXF6BJ7gaxxBDsIkXpANZ44l2xPKyDD2dSqQGle+/JWzt6l5G81ivLAmTih5U7PJIErTEff+q75DoCD3HIWKnBlkRreNwrgCaMawkpEWrlwDWRIdsGSXIlmPTfNjZCjNjL5uqN+pz+xFBkY+4TKy4rv4ApmXfVFUrY5QDaHYGgunx6BeJ408TY63GxYSqw0JfSTr0YYYshbWmnfa5g0UUjJ2x0JWqhos8CBe2Lr7TfqY/uNOgaPt9/uH/9g0hwO+Kr5dqeI3DdAmPuGKD+1HfMJnfOm6u3j7801ecN3F7fWMW1f3JA7eq+Ai0wTjspritX3GC+61Xo/N7wAAAP//cXedB20CAAA=")
	data["docstring/def.md"] = decompress("H4sIAAAAAAAA/1ySzYrbQBCE73qKAl/sZSOTzRPsIYdAjoEcQg4tqeQZmB9luse28vRhxoaFXBrRP9VfteaA48IVSSKx5hJPmHxaFNJTuslMMFnZh58+hF6E4CqhEpafbfAJ5oi5lsJkH6OvuDk/O3jFuSrLGdOOhavUYCPen4I+XRRSCB9jNZkCIWlBodZgTVsSWEoumOjTBUW8coFfW0HMGDdrK6IsHarwU+c0xzjih/OKiU6uPpfWtvh1Zec0Jwkxq+G7101fIdqGdtya1wsTi4SwYxUfoD4wWdgbkNbZYRaljsNwOOA94etd4hY4DHhc9D4AwDHK1j+A45rwa/+NHS8veDs9s8o/n1t4a+HL6TQMnZcPtQfIXChGCIL83RFlw82xEJTZgYGxeclr/wXmComt5KtfuDTRyjRTu/Fcp8AFdcsJpVXURnyzxxZzTHjeTexjl+Uu/N9zwPl+Hod/AQAA//81Ap3HPgIAAA==")
	data["docstring/defmacro.md"] = decompress("H4sIAAAAAAAA/1SQwY7iMBBE7/mKEkgbw65gua60QhzmG+aAODR2R7Hk2Bl3Z2D+fmQTGMghqtSrVLd6CeO4G8jmhEgDY2ClPY5Fr9cndCkPv1c4++gEhMzkOKPmm3cfQiWgm4Pz161EE7Rn2Clnjlo9GcnyH1x6b3t4wXYSztvyh+OOpqAbHOYWL+DrSNGxK7w0zXN9RMpFaAIFLaJnSJqyZdjkuBU40rLnmFk4KqlPEWfuUmZ4rd2fFCZSdpumWS5xiHi70jAGbho8XcOm6BoAWPgCB44qdVxbSFsvs6iB4y/YQJOwnOq3ufQcYYQ/9newqgAwvoP5jx1M4PiAdwqY2cLfJy94UbQyTqqc//nuQV7yP3q3ekmkKLedYTKLzu/H7Po03wEAAP//H/DT9wgCAAA=")
	data["docstring/defn.md"] = decompress("H4sIAAAAAAAA/2yQu24bMRBFe37FhdXswnpYKQ0EQYrkCwKkMFzMcofeQcjhmg84cpB/D0g5lgqzIe/cefFsMMzsFEqBEbjQFzxQesqPcDGF2xGT6JxBPSGvZBmuqi0S1fwU77sPeg9iOp17lYiyMGxNibVcyrd4WcQukIxDzZwOrWJmR9WXvTGbDb4qvv2msHo2Bm/rOZkMAPy5n6PFTeJSk2aQgj2HNiC6Ps/JFJWsFWR+rqyWb/72ygd57Pdgo879BQyfIbgbcXetjyOO1/rTRd+zz9zitxicTBh23R+v1HFsx5gfi2Tw+RtYOTWaGYltTblhepGyQCMKiYcl7xHXIkFeqVHcgnRuduAQ32J7fI8JhBDTe0fS0jNzIftr55Kwzv70AQQ8sXLqff5vtUVm7szmaGtjeLajw8HT62mX+fmwN/8CAAD//z9A8NUhAgAA")
	data["docstring/do.md"] = decompress("H4sIAAAAAAAA/zzOv8oCMRAE8H6fYuCaJHx8aGkjWPgM1kH3vIW9vZA/mscXI1y17PymmAnusWHe8hqCB7+itli5YG1aJSkPKnQT1V3B8b4MgBhqy/aHzN8r9kRdGLNY3PuyGWKB1ILMpWn9J5omXAzXHtekTAS498IGd0bHyROAsculbOi/f2QhoON48N7TJwAA//89L5iPugAAAA==")
	data["docstring/drop.md"] = decompress("H4sIAAAAAAAA/1SQvU7sMBSEez/FSFvcWBdlBfsDLQUFPRIFQitjT5RIjh3sE0h4euTNArvtmdE3c2aFyqU4wMYxCDLf/2uUQ4a0RNOlLKBnzyAZsSmOkcEyq+fOeyTKmAIMvPmaf0VIawScrB8dz0nrY8z6gljUIcWPztH94VHZGKwRBiN0usbj4sxxTJZnPpOI3MYkTCU2HF2pyFnolr+uYALYDzLDd1nwWaq/8dSerlZqtcJ9wMNk+sFTKaBybDDhX3WNG2yw1frnOONlhz1ucfe63Mp+G0yYtVJPbZfBBbPknCYqtS5HOlRb7LDXh1p9BwAA///EfzvxhgEAAA==")
	data["docstring/eq.md"] = decompress("H4sIAAAAAAAA/3SQsW4jMQxEe33FnN3YOGP9DVdcESBAypQGrR1ZQmTKFrlx8vfBeh0gTRoSGIAzj7PGhlek1s/38XcLp7mhJAiMjpbwLnWiQTpRRqqXKBXe4JlIpZuH11IrOn3qikOSajxADNaazrs4qLFN6uwGWeI8i6MYtPkvtgNePLPfihG3nwneJx6GENZr/FP8/5DzpTIEANiMTMhYZdbaVttF4/VbQN6G8FzeiHMzx6VzLFGctoPnYkiTRi9NEUUh1RqOhPIkzhHHz/ngQh2Lnu6Y+z97xCxdorMPeNLFJYpx98BcyI98wHOcq51/llrvRSzF8jpJHcJXAAAA///qslI7jwEAAA==")
	data["docstring/eval.md"] = decompress("H4sIAAAAAAAA/1JW0EgtS8xRSMsvytVUSK0oSMxLKVZIzEtRAAmXJpakFiskgmW5AAEAAP//Q1NDtisAAAA=")
	data["docstring/filter.md"] = decompress("H4sIAAAAAAAA/1yQwWrzQAyE7/sUA7nY/D+BJD32UkoPvfcWgtk6ci1QtO6uNon79GU31DQ9SqOZT8wKzcBiFDFk7ZHo818L8V8sM25CKstM2lNyz5G8UYIvJ/Mi4DKGROiDGqmBE2wkREpZDGGAnyaZWT/qeorhzEc6VqBxUFioAgmdSC0Vx93hwl/jdfgTXaYSz72vWYUds40zGg2GbvCSqPuPOihL1xaLVt/ZSyZcWATvBNZecsGx/mKUr3/4a+dWKzwpXq7+NAk5h6W9ZlDsrwc0j7hi17bYb7DFDg+H1rm3kdONE8lyvOXfN9g1G2zbzn0HAAD//zc5m3CQAQAA")
	data["docstring/first.md"] = decompress("H4sIAAAAAAAA/3yQP0+EQBTE+/0Uk1whJAYSvRipjIWFvf0dLLPy4rIL+8fDb28Ao7nGdibze2/mgMJIiAmRc4nAlIOLSAOxy7Qc6RK82cTIOdNpqrdBIkx2Ool3uIi1P+F/shO1GGH/S7mFD3BiIdd0SATHKX1VeE24+Gx7dERHRyNaWovkoQfqDxgf0OKztfKHRY7i3nEuIuenrdgZHY0PhG6tXb16+7Be79eBMdWVUocDnh1elnacLJUCip4GC26KpsHDEfd3OD6W5Wrs/ZZSKWxDcA9d7XBqmlOlvgMAAP//i7WKe2ABAAA=")
	data["docstring/fn.md"] = decompress("H4sIAAAAAAAA/2yQwa7aMBBF9/6KK1g0oQKVlKLuEIv+QaUuoghNnAlYsseRx6bw908k7N7b2kdHc88a1SgQCnxC4EwntJSu2mGMKXyvYRNTZgUJSKI8QyyKsYjNLor557x/I18TyDfKCPREz5hIlQdQikUGOAFhdEnz1npSRSARTjtj1mucBX8eFCbPxgDVwCOGWHrPBgAqzxltKB5NNz9gXrFaEMWdfGFdoX10qDYbPBCKr+u6nmU5bu9sc0yLK9D0duNbtUeDnzjgF44z/vfmFLxcgv+vtYlzSa9djMWCS9vggCN+Y/8D+6a77Iw5fyqhmFIcimUQrI9aEi9xbJwc6yzsnQxOrrp83OjO6JkFiUdOLJYHjCmGmdWS5o5OrlAbJ96ZjwAAAP//Jcru6MwBAAA=")
	data["docstring/future.md"] = decompress("H4sIAAAAAAAA/1yRUZLbIBBE/zlFl/0jqRLvGfKRC6T2AIuhtaKCBxcM2tXtUyDFWYfPGabfTPcZw1y1ZmJO+TZNI7jaWK2ywKJQkebeKrBlE7fkJKmWuJlf1Jql/ToEgkCXXacNtbo4DUkueF2Ie05r8PSH2keIEVc+cL7NN+LdZquELpnWNyF+0tWm8w1WPKxscDbGAk078MDsktM1Jvd7QhUN8bFQwWJX4koKXLrdI5Vx+we/GHM+44fg56dtXWOAwXNut/11yADAoOn7SqcpY3insO3aG8DAW1CcMv1pfC6lbOWd/1c3xpg+TmN7nTdXHY15XUIB9zWaab46fnG5nU6Pl7nqC3SxCpdkZdbSj80sNWrpAchTZOjGIIgmWOw37Mm8dfJblz5y6V+/WLiz8yPy5hsv5k8AAAD//3S7uL4+AgAA")
	data["docstring/generate.md"] = decompress("H4sIAAAAAAAA/1SSTXLbMAyF9zrFm3hjpYlyhi5ygU6mm04ngUnI4pQCVP44UU/fISk5zlKA8N6HBx5wPLNwoMQYNczfeuzfEYTIfzOLYVBcxUxBRXP0a/d8IZ/rP2lixIWNGx3bKhHhpI4uVGXTFJgsdAR/sMnJqQz4wSkH+WKRJkp4d97DpQakFw4gWcts8Smm1ZISKDB4dimxHfCzNW5qyNHJGQSvhrxfEY0uBTCLKQS7ZAHG27FMNfn+bcDL1tgUKzVbUME1+hg0Jyf8gJlJiksDQjFC0orQFrk/eTV/7pElOV/9jIbAcVGxZZA9zywJLiJwVH9hi9NaXSTmmcOOuYc0dN3hgO+C5w+aF89dBxwtjzDqNcTPY3YA0Na6C2zv+tuCBpIzf62t7L2+3/V9XzWTPl7YJA2bct91L5OL4Obb1js5sRXP07/185Ch3ratUrrXB7YHVIpCM+OpiT+1yG88y4/NogR4zY/KkYKjk+eS2ZjLZbes7AOo4chG0N5MU3z9VWO47n5d+Pfr0P0PAAD//3asAtAFAwAA")
	data["docstring/gensym.md"] = decompress("H4sIAAAAAAAA/0SQsQ7aQAyG93uKX8lQIgEPUIaqQ4fu3SoEhjg5S5c7OPtS8vbVNYGOlv19tv8Wu5GjLhPU8rcO98xkrCCUKM/C0GW6pbBHUR5KgERMdM9J3c8BVCGJI0TxyGmWnvs9zJO9G38kBNy40j0s4VkoyLDAPG8LwoKRI2cy7rdlR/zyohhKvJuk+FbrRvWcw1LlN/Y0S8oYUoYn9QcjCR+N1mN1iUavg5fRBxm9VW594Ohc2+J7xI8XTY/AzgG7wIbf+smkmSk33dkBtSdq+KKPYsb56zbZeA4hNWdo11XD6VTD4GeRmQJHgyUHXK+reabc/kdq1Tn371lej1gDy2wlR1y2ycvR/Q0AAP//Kne9pacBAAA=")
	data["docstring/get.md"] = decompress("H4sIAAAAAAAA/0SQvW7jMBCEez7FQC7OBg6u7ho3QQKkSBukN1biyCJik8ruyj8I8u6BGBkpuTM7/HZWWB/o6Iue8M4bInuZjv6wgdI18UyD4CzHiWhvsyW80ifNBh+4CJfkQ8oQGD8m5o7wQRzJIGalS+KM1VR3bGSX+sQ4p23x0tdp/bzQkIuD12R+j607S/Df+ZXB5AO1KgvwQpJs5p40M6IoJIOqRetckjFuQ1it8JjxfJXTeGQIwDqyh5aW6vjcZTkRzVNpG+zkQPz7/7WZTXNNi2knMSrN0Fy05EOzCeFtSAb+ZC4Ihv2i79Gyk8lYie+dpDPrxfmPoyvZZa4w/4aPWkaq37bhOwAA//+YoO0YpQEAAA==")
	data["docstring/go.md"] = decompress("H4sIAAAAAAAA/zyNTY7rIBgE95yin72BLHKHLN4N5gJfTNugwXwWPxl8+5EnUralqu4ZdlOsWvbbzUHqmZdQNGuv6QRfkro0VgieSZdv8xWIo+grevq/quInpoQnP7JHzBBUHlKkES0Uioeu4ODSW9R8xyOfKKw9tZg3XCUvowW+jz6rPtZFiqe/GzPPeGT8H7IficYA1nPFgD2K7rHSuYttCjswBaak0xvVVmCHw4RT+7/Jmd8AAAD//0iWKE72AAAA")
	data["docstring/has-meta.md"] = decompress("H4sIAAAAAAAA/5SRQa/TMBCE7/kV03KgKdAK/gBnJI69IRRt4k1t4awj77ot/x45Ka9PT728a3ZmMt/4A3YTG33HmPL0qYWxmuLq2TxnmGfMOV2CY7coFJ4ujGpxZNT8GEEx3k98oVjIGJZAIsnI2KF+ZP1cswTmg2IsMlhIgmuIEZmtZEFnuXB3wMkzxpDVIEm+PGLuWi3REGRp9pKzRgQ5oxspKneHpgF2jkeBS6WPjF+339jt97jhW9vW40q9XtumOdVifKNpjvykF3oeqOjjn4qBBD0/QA9N8zP8YUxJDXNmFwayFfw1c7VR1FS9wueFrf9bDTOLqwwV7bg5YvCUaTDOdZSgmJhEYZ5s2TyNz55nKstyiyRd/7cLSdZJNiv21nOMaYuv7yBXy0HOChIHKVPPedngo71Z4V8AAAD//wftduxSAgAA")
	data["docstring/if.md"] = decompress("H4sIAAAAAAAA/3SQu27rMAyGdz3Fj2Q4FmAco01vU4sOHfoUgRxRMQGaCiQ5cd6+kI02XToS5Mf/skXDAadEHmUgBUmmN4sTpRDTmJF5PAmhT04PA+vRfIZ6CDo7mVwhv7B8cIXAGSVNZbii0ViwD04y7Vssg7LsbbuwXVXqUAUqc3vl1CNRmZKSbxHLQOnCmVaoOluhFhzg9NriwiLo6Y8X/43ZbvGu+JhdTWEM0HgKmPGvucM9dnjAI57wjBdrly0HNK9ohBSzxc4aANjM1WbPx82vMY9OZGPNdyH5OvZR0E0qlHNXT6ZMHqy5kPOIAR2Hrl1rroTEIx+c3Lr9icN6plRqgK8AAAD//9J6mECfAQAA")
	data["docstring/is-assoc.md"] = decompress("H4sIAAAAAAAA/4yRsc7UMBCEez/F/LmGXxw5AR0NoqBAorwOoWjjbGILx468mxwndO+O7BwnQBSUlndmZ7854AWJJPseY8rzy2coiwoujtVxhjrGktPmBx7qhIAyo0o8qd9YzKcRFML9lzcKKylDEyjug8diE6HOC8Y1WvUp4uJDQGZdc0SneeWuxdkxRp9FEVN8VbW/5mQNCh9roIfHLvdxQjdSEO5aYw4HfIj4+J3mJbAxeNz3412kmdH0PoTm9nguPE3X5oYvr/EGb78+G3MuOXk3+DPmvgQ9W1qFaxZ1Pu9k4AWEja2m3Brz2X9jzEkUS+bBW1KW418MbCEUJKFnRJ5IeUB/LYKF41DuKitOTydYR5msci6QvGBmigJ1pJV9Gv/V1LxWklrsf2+srVie7lwaxyGk5r8A7D2ZnwEAAP//8byYiTUCAAA=")
	data["docstring/is-atom.md"] = decompress("H4sIAAAAAAAA/5SRP48TMRDFe3+Kd5eCRIf2xN+CBl1BR0mHUDSxZ+MRXnvlGWfJt0feJamgoJ0/b95v3g57sjJ9xljq9HSAsZpiiWyRKywy5louEjisEwqqjL4h3r2sJYjCl6wSuHKAFZxuE5ARYvCUc7FeHltddflCqZFxAOWApbQUUHpnEeV7t2uJKadxcG63w0vGl180zYmdw813loRPkVMqeOwK/Hhw7lsUBW+zWCQlVLZWM45WGx9xYk+tXyIf0a9xp9hMD859lZ+MqahhrhzEk7G+hnXRsWVvUnKHAiVdaTOfV5jTtS/MnIPk8/q854dn+EiVvHEdsPqamLLCIhkoJZTxb2+emhr+fO3uC9g/bNSv9k94g7d4d8D39/iAjz/+g9qXaS4tB9h1ZoU2H0GKJD36HsiFvZW6Zd09/CPewf0OAAD//zgU0jU+AgAA")
	data["docstring/is-indexed.md"] = decompress("H4sIAAAAAAAA/3ySzW4UMRCE736KSnIgK8JG/JMT4sABiWNuCK16PTU7Fp724O7ZhLdHngkBhcC9u1zf1z7DedKOt+zeoy91fLqB09xwM9AHVvhATLUcU8dumTBIJe6WYPw+UyMtfOohOd+N8Ch5Fie8QHCUnH6PwgdRRFHsCYmRZuyw/7FmXrQXFT4kQz9r9FQUNylnVPpcFTuvM3dbXA9En6o5tOizh31+rdicHUkXjPu4NSnpAbtesnG3DeHsDB8UH29lnDJDwB9enpw/xwu8xKsNvrzGG7zFu6+bEK5bRa4LjzUM4XP6RozFHFNll6I47eIBWvMg2UqToTyIry6myonatY6t+uXJJeIgVaKzNvZkGClqTaYv4kv/2K3GeRHkLX49w1+X2y60J/e4pwNzLqe4uvov5GoOe0aZjTCvSQ//+Bzb8DMAAP//ywxy6GgCAAA=")
	data["docstring/is-len.md"] = decompress("H4sIAAAAAAAA/3yST2/UMBDF7/4Ur90DXVFS8Z+eEAcOSBx7QyiadSaxhTMOnvEWvj1yUkBalt7fjN/vN97hKrG8x5jL/HQPYzXFfWALXGCBsZR8jAMPa0JBheFzFeMByt8ri2d1n0ZQSg8RPlKqZAzLIBwpxb9RWCCBJ0HhJRdDNEVimSxct+cEFqJirOItZsF9TAmFrRZBb6Vy3+EuMMZY1CBZnp2W+T2iNRmirAx/1m2bokzoR0rKfefcbocPgo8/aF4SO4cHIU+unuMFXuLVHl9e4w3e4t3XvXN3rR5v4XPtnPscvzHmrIal8BA9Gev1CVYTQEkzDgzhiRrA4WcbWFiG1q/Vvrm4gQ9UyBuXxh0VM5Nos2ir8TyeO9JcVznW1m/+/zlZt5JerKiXgVPKl7i9fRRwM4YDe6rKUCtRpv/8iM79CgAA///IMqaLWQIAAA==")
	data["docstring/is-list.md"] = decompress("H4sIAAAAAAAA/3yRvW4bMRCEez7FSCoiIcoJ+Y+qIEWKACnVCcFhxdsTifDIA3dPst/eIM8yYMNwS8ws55tZYR286E/0KQ/vN1AWFVwdq+MMdYwxp4vvuKsKAWVGcYj504NCeHzmC4WJlKEJVAXb4o5Q5wX9FK36FHH1ISCzTjmi1Txx2+DgGL3PoogpfijWm0ymoPCxxng6Mbt9PKPtKQi3jTGrFX5F/L6jYQxsDG5Q79Yf8Qmf8WWD41d8w3f8+Lcx5lAy8ax+Hmm+iBNbmoTrx8I2xZkeXkC4sNWUG2P++v+MIYlizNx5S8qyfQFsKYKCJJwYkc+k3OF0Xwwjx65QlD92ix2so0xWOZdGvGBgigJ1pLXm1L+2xjDV2rScr6s0lX4x4y8dh5CWOO73b2PPS5iHAAAA//9jO4yZDQIAAA==")
	data["docstring/is-mapped.md"] = decompress("H4sIAAAAAAAA/4yRv87UMBDEez/FfLmGTxw5AR0NoqBAorwOoWiTTGILx7bszR0ndO+O8odDIApKW7uzM7854MUkKbF/jyHm6eUzlEULrpZqmaGWSDleXM9+nSiQTGw75tMA8X7/50X8LEpohOwT0FvicVEJUOsKhjl06mLA1XmPTJ1zQKN5ZlPjbInB5aIIMbzaFfbBMnuFC6uhh8i278KIZhBf2NTGHA74EPDxu0zJ0xj8DvjjXZCJqFrnfXV/PBPH8Vbd8eU13uDt12djzotTbgp/Gt2uoGUnc+FqRq3LGxq4AsGFncZcG/PZfSOmWBQps3edKMvxLwqdBIgvES0ROIqyR3tbFhJDvwRbTpyeTuisZOmUecHkCiZKKFArulYQh39VNc0rS13kNwr1SuTpF5LK0vtY/Vf2raSfAQAA//+urcIYMQIAAA==")
	data["docstring/is-nil.md"] = decompress("H4sIAAAAAAAA/3yQva7TQBCF+32Kc28KbsTFET9NKkRBgUSZDiFrYo+zI8az1u44gbdHuw5IIES7O3Pm+84OTyb6HlPK88s9nIsX3CJ75AyPjCWnq4w8tokCygwTDZ8mkOr9ka+kKznDU/18rosGj1IwrTa4JMNNVJHZ12zoPa/cdzhFxiS5OCzZKxP9NVVWdYg1gN8J27LYBf1EWrjvQtjt8MHw8TvNi3IIuNu8eHqNN3iLd/vKsw/hVFl4G/sTZYvCmQdaC28XG1JVgxQQVIp3IXyWb4w5FceSeZSBnMvzX5YDGUhLwplhfCHnEecfdWFhGyt7PXB4OGCIlGlwzrUGKZiZrMAjeSs2Tf9qf15bV97iRbtm/NCUHyOrpkd8OR6//ld46z78DAAA//+Arx2/+QEAAA==")
	data["docstring/is-promise.md"] = decompress("H4sIAAAAAAAA/4xRvW7yQBDs7ykGaLA+PiPIA0QpUkRKSRdF1mKvuVPOd9btGpK3j84GS0Ep0u7uzM7PCus+xc4JP6KNqftXQFlUcLGslhPUMvoUz67hZrwQUBpHGSTmpQV5f93wmfxAytAIut1sMkeAWidoh1CriwEX5z0S65ACKk0DVyUOltG6JIoQw/8r+nYpg1e4MOqZWSYCF06oWvLCVWnMaoWngOdP6nrPxgDrhlv0u9loUczD/d1wjqLf5eXbDns8vBfGHLJ4njh/ap/+4sg1DcKjPLUuTVnBCQhnrjWm0phX98Hooij6xI2rSVk2d8HUFEBeIo6MwCdSbnD8yoCeQ5Ot5hfbxRa1pUS1csrJOUHHFARqScdGYvtbd90wxquZ/tZhOaa0mM0vLXsfl3+yP1VnvgMAAP//dTd9zkcCAAA=")
	data["docstring/is-seq.md"] = decompress("H4sIAAAAAAAA/4yQTU8yQRCE7/MrCji8kBchfsvJePBg4pGbMZtmt5eZODuzTPeC/nszy0ei4eC5qyvPUyOMhTePqGNq/k+gLCrYWVbLCWoZbYpbV3HVJwSUGMKbjkPJYl5qkPeHE2/Jd6QMjSBsybvqFJ3mrgC1TlB3oVQXA3bOeyTWLgUUmjouZlhaRu2SKEIMF8f3Y1Q6r3ChBzvV7BtcWKOoyQsXM2NGIzwFPH9S03o2BgfLf+NLXOEaNxO83eIO93h4nxizzFi8D5+jMubVfTCaKIo2ceVKUpbpL52SAshLxIoReE3KFVZf+aHlUGW+jD0fzFFaSlQqp+zrBA1TEKgl7eeM9bnlm64fRXP9z3Fl1hsOesWhZe/jEIvFX8S+AwAA///C9CJE/wEAAA==")
	data["docstring/is-str.md"] = decompress("H4sIAAAAAAAA/3yRQU/jMBSE7/4V0/awjXY31S5cekIcOCBx7A2h6DV5qS0cO/J7aeHfIzstKhLi7JnxfPNWWIumO/QxDb8rKIsKTpbVcoJaxpji0XXcFYWAEkM0uXAQ89iDvD8/8JH8RMrQeBH8yQEBap2gn0KrLgacnPdIrFMKaDRN3NTYWUbvkihCDH9n90Uok1e4ULp8hsz+rGp68sJNbcxqhfuAhzcaRs/G4Az2a/0P/3GD2wpLy97HZWXMLlfiWfq10RyHPbc0Cc+/lmaZEk5A8E60NubJvTKGKIoxcedaUi7E17AtBZCXiD0j8IGUO+zfs2Hk0OX++YPNYoPWUqJWOeU1nGBgCgK1pGXj2H93jGEqk2mOP29eF/LFNXqF5+325Ufq+Q7mIwAA//8Gxb6fDQIAAA==")
	data["docstring/is-vector.md"] = decompress("H4sIAAAAAAAA/5SRvY7bMBCEez7F2C5iI46M/MdVkCJFgJTujIOwplYmcRQpcFf23dsfKNn3h2uu5sxwvtkFlie2mvJvtCl3H1dQFhWcHavjDHWMPqeTb7gZFQLKjMkj5l8LCuHywCcKAylDE+giWZeECHVe0A7Rqk8RZx8CMuuQI2rNA9cVdo7R+iyKmOKnyXwVyhAUPo5lHkMmv49H1C0F4boyZrHAn4i/d9T1gY3BE9yH5Wd8wVd8W2H/HT/wE79uVsbsSi+e9C9rTZk4sKVBePp6rFdY4QWE4EUrY/77W0aXRNFnbrwlZVm/IrYUQUESDozIR1JucLgvhp5jUyDKB5vZBtZRJqucyyRe0DFFgTrScenUvnWSbhh30xJ/OU014s+u/HPHIaQ59tvtu7mFbYrNM/AptDIPAQAA///GzZ6GPQIAAA==")
	data["docstring/last.md"] = decompress("H4sIAAAAAAAA/3SPsUrEQBRF+/mKC1uYgGRBg+xWYmFhb787mdwxD19mkszEjX8viaJsYfsu5/DODoXalJE4lpiY5ykk5I7YrlT2DBnRb7fEcWZwNK+dJPg5uCwx4CKqP+z/6EAnXtj+Sm4RJwRRyLUcksB+yJ8VXjIucdYWDdEw0IsTq8gRrqN7h48TLD6syp8Wc5LwhnOROD5uWWc09HEinFVdt/364L4yZrfDU8DzYvtBaQxQtPRYcFMcj3iocX+H+lCW67A1LaUx2Nr5zVyln+rDqTJfAQAA//8Cb6VGUQEAAA==")
	data["docstring/lazy-seq.md"] = decompress("H4sIAAAAAAAA/0yPsU7EMBBEe3/FSGnsQ4eC6Ogo+AO6iGLtrBULx87ZG5S7r0f26QTTzr6Z2QE60u16rnyBz2U9nQy2kufdcQWh8mXn5BiykCBU8A/FnYRnRLqFeFVqGPCe8HHQukV+UwrQM/sEH2xPnb4UAOjIgskH2wpXOxMmgr17//U3R7uc2gbdIAv9BII1TQ+qOyNejDG91+dyZnILpgNa6JvxOo79qgU+OL2VkCQmHA37XNpX9/XoToUsDB9Klc77YHMi5wLSvlou9Vn9BgAA//8aeNH1NgEAAA==")
	data["docstring/len.md"] = decompress("H4sIAAAAAAAA/1SPwWrzMBCE73qKAR/+GILhb/MCPfTQe6HHIKvjSCDvOtKKNH36YjtQelx2vm93OhwyBZXXHoXWilRYJDLlYhE6bVPltVEC3UfK+ZHbFtLmkWWNMXOmWEUSeARtYn7Mv+iAt78upIrsv1O+I+i8NOPnEb7eJcSioq0eoQVqkeWWKpEk+GVT6oSRSS77lRWzmCqmJsGSCm7bk36FvIClaBmc6zq8CF6//LxkOoe9+L/DfzzhGae+d+591XBPPCx71fPpPLifAAAA///Twvx/KwEAAA==")
	data["docstring/let.md"] = decompress("H4sIAAAAAAAA/1yQP2/jMAzFd32KB2S4GDnYuP83FAU6dOjSqUCHIIMiM7UAhUxNOrb76QvLTtBWiyCK78fHt8I6kWF7w/5IOEh7vN3s8r0psI9cK5IEn3D2qSN1zzElhJa8ETyY+uVbg5zoe1ZEfoE1BD1RiIdI9aKFCazxNvdiP2KaWeLB0E9Ua4hBU+8E/0yYDCn6aE3kjxDPNVqyruUsaEm7ZJBDfiWvdgFG4dK51Qp3jPvBH0+JnMOy/IBv6x/4iV/4XTgsZ8T2D/7iH/7vdrm4DsLBGwaMReHcUxMVNKPmBa6xpKiG4FOiGtVQZZceZwom7bU+Vl/dJ/82Yh5CnC3Pm4gSlF474kBa4lFyPN6yyHrJOSpqAYuBhmm6dKaxpksUVSKrcoqlew8AAP//tr4OOfMBAAA=")
	data["docstring/list.md"] = decompress("H4sIAAAAAAAA/0zMO2rFMBCF4V6rONjN9SVkDymyhtQiOkaCkWRG49fug1+Q9p8zX4+XpGYYq+b3e8Cv0hsbPApXHCf3k0Tu/i9jjbURFGYWa/BKWCS4eJm9MZxkw6R1SYHhA1WhtFnLtcuT7ZeURpR6zw/mefl0ru/xVfC9+TwJnQNegSM2dJEitRuesqOzSOVVTnXDPri/AAAA///Sf+9X3wAAAA==")
	data["docstring/map.md"] = decompress("H4sIAAAAAAAA/1SPwWrDMBBE7/qKAV+sFAJ1Q+m1lP5BbyGEJVnXgvVK1cqN3a8vUcDg67xl9k2DdqSEftILjH+ePIT+giwYKdk9mVgvbO4jMxU20J0vK8BtiMZg4ZG1GCgzysDIbJMUxB6UkixBv2uccvwNV77WfyVERYkVrAWx3x6uBnvnmgbvis+ZxiTsHB7qba84zie0ux1mdN7j+IwOLzicvHNfQzDcgggylylrbd9OOLcdDnjFmz/v3X8AAAD//4F0fMMQAQAA")
	data["docstring/meta.md"] = decompress("H4sIAAAAAAAA/2yPsU7EMBBE+/2KUdIkp4iCMh0FBT0dQqe9eEMsOetTvOaCEP+OfE4HnTVvxs9u0a1ijDlua49NLG+aYIugxI4rQi3Qy3xH5QyfwKrR2MQNJda/M1vYavvmQ8BFDoM4cJmDU4qTZ/OfgmRbnixv8kDUtnhSPO+8XoMQAZ2TWeFivgTB2/6O7nTCjse+L/D+hQp7otfFpyqstuPNIcSb14//nSPRGfge7esqaOask/mozYBReZXj7gGjixOa5udMvwEAAP//4WKBkTkBAAA=")
	data["docstring/not.md"] = decompress("H4sIAAAAAAAA/1yOQWrDQAxF93OKT7xpoCRn6KI3KHQ5TBK5I1BGQdI4ze2L7YLbbKX//n8DXpoGRrXrHqJffC4iD3CbyMIRlRDWo3Ijd+i4XG6mE1/osmDpk0VgFN0a8ljEKYOfgsc5eQT72vZ4hUYlu7MT7n/5sE75kNIw4K3h/btcb0IpYdXcVRLR3T6lj8r+j5znTqpCZbM40bl0p+U3FemE/NuQN5VD+gkAAP//IPscxwQBAAA=")
	data["docstring/ns-put.md"] = decompress("H4sIAAAAAAAA/0yPMW7rMBBEe55iYDVWYesMLv4NPpDSWIkraQFySZArJbl9INpx3HL4duZ1OGu95M2gFUqRMacSe4yivoLaU800MVitfLsPCaGFIOwUNoal5zeIwlZGzTzJLOz/4CtuT0x0qaDCkBg3ozEwSD0K1y3YcYEUXEoqGFl0QSGp7CHzEZAZx2yQiki+VRe+tDW2crw613W4Kf59UcyBncPL7qwVU9JZlh6sO065JL9NJklPvXP/V6ngB4bPl2PTsXIMub8T96P6CJv3wLoPv/bDo2V4k3c/AQAA///Pb6A6ZQEAAA==")
	data["docstring/ns.md"] = decompress("H4sIAAAAAAAA/0zPQU7EMAwF0H1O8aVuWmmGKyAWLDgBC8Qibf5ApNSJbA9Tbo8IQ9Xtt/2+PGAUg8SVjxOUflUxxB5YiwvDay7lPjjmWKo4N8ccjQlVkP3PecDLBf7Jw26qNEh1xKKM6RvcsvkJ2XH75WdiUUZnwsxLVWJmlo97LVMnpXYR2dC0fuXEdOo9y1WV4oe+f3S/D2EY8CR43uLaCkMAxkLHm/T3ra48C2/nnZjeAwCMTQUyTeEnAAD//wJiTz4nAQAA")
	data["docstring/nth.md"] = decompress("H4sIAAAAAAAA/3SSMY/bMAyFd/2Kh8vQBLjm0Dadixs6BN3a7gfaomMCMuWKku/SX19IdjL1RvE9kY+ftMNe84ghpgmint/geaAS8rcDEuckvLCBsFAojO66mtxPziWpIY+8SXmkjJ4UHWOIRT0oN9lm7mUQ9lv/OECywfhPYe35iPPQfKsqhlhyNdVaVxvZ7XS78lhPCpY8cmrKlnmLIlajl6TsERNIwSnF1Ookxv6IH8wzRDGJ+jV6G19XTYy/nOLHjqrTud0Oz4rvbzTNgZ0D9p4HBHzYf8JnfMHpcKjFSjHghIfXFPXycHDu9ygGXu9tgQwvm/6Cjnsqdtv7hH1d5Gse18Khpu34Gls+Bqu/Y7gDDWJ5TXhu6T1+bYjwPM9BesoS1bmzWmZqDYqJXlqbJ83jE4aifTU9bgD8nbLBSj+CrI0xkHos3OeYrL0zBYsYaWFQupSJtXrqVPbwkrjP4Yoc66zp+D63/zB7lRBwFQ7r6kZT5Wf1gWn9cnPiRWK54z26fwEAAP//JdYY4cgCAAA=")
	data["docstring/or.md"] = decompress("H4sIAAAAAAAA/3SPsW7rMAxFd33FRbIkeXkO2gIdumXI3LGjINh0JEAWDYqO078vZDuIO1STQNzLc7jFjgUtS3c47NGTlG+GQ/Ys+r8OUg9BQ7rixHIyl5uLg1PKUE+Ys61wh0itQhkSrl4rnDMyc4LL4ESgZ43hoDKo/0YZ0hFjiBFCOkiCeqfzvMKnepIxZELQOdQL10RNWbJsLGLFJNFdJ53KmO0W54TL3XV9JGMwXbj7hxe84m1vML3WxUzLf5NYFwFqNntjvlZG9t0ekehGsmbaqW/hUgP7u2+PGD0JufzxgK9hKcQH9i9kOSirTJxnyFbmJwAA///Ke/TyrQEAAA==")
	data["docstring/partial.md"] = decompress("H4sIAAAAAAAA/0yQy2oDMQxF9/6KC9lkSBsaKIUsu2g/oBS6TDSx4hF1ZONHpvn7Ms5za52je60Z5pFSEfLYV92Bklt0cKycqHAGteciQdFTZougIFwVitGfzBeXmnRilUd8XvlxCJkhKmc0uXpgLRmUGDHxcx+qWpSA0sCYwlEs2yV+BlaUgcp9l2SIHsMv2yeQnm7ww9pRvEeWQ/Qn9DxVY52IFvCYeFd63ofE2JH3oq5hIYkTJX+LXhozm+Fd8fFHh+jZGGBueY/oa1693M+3wCveum4aX0brzpjvQTL4rJ4rpnYtbFbrDShPn69uwHbefKy7LUa+lGK7NP8BAAD//+QMK6ahAQAA")
	data["docstring/partition.md"] = decompress("H4sIAAAAAAAA/0yPzU7DMBCE736KkXKxJWia0tJyQhx4AyQOCFVpsmlXctbGXvPz9ihBLT3Zmm/0aaeCjW1SVg6CLhRRZKX4iEwfDheU0U5JIenIvLL3/+iKgEUDjimUmBEG1LOwBnkaSTTfgKVL85/liMMP9ESQMh4oTf1zDz0NLNSDBfV0Tg0b0kXHwznlDAmKmMIn99S7hTFVhSfB83c7Rk/GAFbDrees10NXuMNbMz9rbHCPLXZ4QLN8d86YlxNn0J8BX9PYRFqSYG9tg5WDXWPjYLfYOdhm6dx+YX4DAAD///gpY+NJAQAA")
	data["docstring/promise.md"] = decompress("H4sIAAAAAAAA/1SQMW7DMAxFd53iI1mSIMgdMnTo2hswEh0TpUlBluz29oXgxkVHiR/vP/KIUy4+ycxn5OKpRZ5BhmaFZ9eFExbSxuGDays2g/CbhxjqyBi8TPABhKFZrOJ2w/sAylmFE1apo7famVSebWKrV9RR5j2OVVRxeajHz8sVK0kVe3YuaCtHdTwYiVUWLl1JCIRIqn3ULV5SO7SOVCEWtaVto1f7Dfc9Hcngpt//6W6RbyEcj7gb3r5oysohAKfEA/Lfwc79L+Mwsqofttc5/AQAAP//etNsWFIBAAA=")
	data["docstring/quote.md"] = decompress("H4sIAAAAAAAA/0zPQU7DMBCF4b1P8aQu2kgoVYELsGDBgh0H6DSZkJHssfFM2ub2KAkgtrbs9387HL6m7Iwh19Sgsk9VDT4yrHAng3C/3kEUPTkh5Z7DO5OKfsJHckQxN5D2sDldcjTcJEZodlwYfKU4kXPf4mMUQ6KuZoiBh4E7lyvHedujxCBDqVxY++V7UvC9VDaTrLiJj8sRlWxecxkZh+P+2LQh7HZ4UbzeKZXIIeBXdTjhEU94bpoQ1vW1bFOuo1GcK8XVgEo+cl1QCq/zCsygUn4KdUoXrjg9LJkyQBw3rgzCMGnnkrUF3nzB2axZ55Qn27KX5/8o5/1f2bkN3wEAAP//N4kNS4UBAAA=")
	data["docstring/range.md"] = decompress("H4sIAAAAAAAA/1yQQWrrQBBE93OKAm8k0NdI/ugA5vMXOUEWIYi21LKHzPQoMy0j5/Qhim1IllVF16NrhyKRnBjBCQKtcDKUGBKTcgZhC82/h/b0cUXm94VlYOiZFHPizKIZembIEo6cMqYUA2xwYlE4GfyS3YVLaIQNtFoUvN7NCscrrJMhcWBRW+PgPSidlrDVUmLEWV0U8vWtc+SJFv/FjOibvrq1/rCdTH0FkhE2K8+/0rav8TQhir8iCj94cBlzihc38ljBbVq390dQ/ubUxux2OAj+rxRmz8YAQKHxz4UHjQmF0huju2/bNnAyoSvL0phn5z0S65JkW+x20r+0DdoO+wb7Dn+b1742nwEAAP//JVxzAJ0BAAA=")
	data["docstring/read.md"] = decompress("H4sIAAAAAAAA/xTKuw2AMBAD0J4pLNHALixxil2kSaL7CMZH6V7xTlwuIyL9xlbA7cXTY6FNCn3kRA19ywZF0NL2rpbliuMPAAD//w7hDINBAAAA")
	data["docstring/reduce.md"] = decompress("H4sIAAAAAAAA/1ySTY/TMBCG7/4Vr9QDrbZELFA+jiBx4L6cEFqZZNKM1rXDzDht/z2yk7SIU+TRTJ5nPjbYCnW5JfQ5tlD687DDHNHyyhRbUvfdSLyRIk0k8FAypP6esJ9rOB5hA7GAAp0omsJSSed4DAQhzcFK0uRDpgZPw8w1ThGjpIk76nDKajD/QrBzgpdjrr9a0lnU4GMHpTbVz+xwR1JsU45GQh28UDECRzb2YQYr/DgGpq7Y2eDtJlEZQr43kn0tnJ1Ls+U1Ck2csqL1oc3BV3FWZC0srTmL4aK9x3ngMEtEutiq+X/V0s1a1ji32eBLxLeLP42BnAO2HfW44Ocj3uId3uPwa7dGr3i1/YCP+ITPeHyzq/Flrw+44Lpz7mlgxZlDgJBliZVax4Hnw+G5ce5HDPxCSDaQgAu1mNQedR3AP8dS4mVGrEiRkGK41q0p2PQ28EC9LZhe0mnpdTmb284bfM0GrlMhr1dQTPk4lP0IjZIqtKJ+0+AnToJcjgo+3kjzkb1e5ztRa0ka9zcAAP//5r+52eICAAA=")
	data["docstring/repl-cls.md"] = decompress("H4sIAAAAAAAA/1JW0EjOKdZUSM5JTSwqVijJSFUoTi5KTc3jcgaJgAWSS4uKUvNKFJLz80pAdH4akjo9hZCMzGKFzGKFRIUg1wAf3fy8nEqFtNK85JLM/Dw9LkAAAAD//yVoAgJhAAAA")
	data["docstring/repl-doc.md"] = decompress("H4sIAAAAAAAA/1SOQQrCMBBF93OKD91Y0N5BsDsXIl5gSCZ0oJ0EJ4Xm9tKKC7fv8x+vwynmgJTfS4+oXmZujpjDuohVrpqNbj9cJ/mf9t9BvUjQpBIP0xmawNYgm3r1Aa9JHepgPMfH/ZJtbkirhd0xEHUdroZx46XMQoRvE3uz0NMnAAD//xQLeSyiAAAA")
	data["docstring/repl-help.md"] = decompress("H4sIAAAAAAAA/2SOS27jMBBE9zpFLWVh4APMbuDxbhaTIAcwJZUkAs0mwyb9uX0gOTFipLf1XnV17UJJuw73++stibuhLN6wJgg0czObrh3jgCnmsMFf4BiHGqjFFR91j+PVhST8jdOGz3F3arq2GqH2+eSwOJ2JoeZMLVAXaMkNfJJXY4p3exB77MNB6DLKQtiQSW269r368gBeqi9b/Hr8/69p3hYi5RhSQU+JF/iRWvzkaRv1Y8UvTFEkXjiiv22I1tAzI07gNWWa+air7AouXgSLOxM9qeDZSXWFI7w+ta9bYHd1jz8wr7PwWx8Gp7DkFKFK8UkI8UrbNx8BAAD//3pr7pSgAQAA")
	data["docstring/repl-quit.md"] = decompress("H4sIAAAAAAAA/1JW0CgszSzRVACRxQolGakKQa4BPlyeubmpKZmJJak5lQqpFTCp4ILSkpLUIrASPYWQjMxihcxihUQwXzc/L6dSIa00L7kkMz9PjwsQAAD//+kUzLVaAAAA")
	data["docstring/repl-use.md"] = decompress("H4sIAAAAAAAA/0TNQQrCQAxG4X1O8cNs7MLeQaQ7FyJeIITUGWiT0mSg3t6FgtvHg6/g1ENhMUAq20sDxqvGxqJ0/ZWsCun7rpYQt9Qj/9eIZ22BFmA8pvvt7La8MXeTbG4jUSm4GKaD121RInzF2X2gTwAAAP//xZr5jn4AAAA=")
	data["docstring/rest.md"] = decompress("H4sIAAAAAAAA/1SQvU7DMBSFdz/FkTqQLIkEFaITYmBgZ28d55hY3NqNf2geH8VFqdgsn/sd3e/u0ESmjMS5RWQu0Sfkiai/wdZ34lzoDdXn5BJs8Sa74HF1In8M9DaEPOkMLkbKyFuXdTFlUHimv5deaJx1HDeyw0fGNRQZMRADPa0zTgtygJlovmFDhMaPFnenUJLzXzg1ifNr9ThhoA2RMFpkzfq6QI8Q0a9efafUboc3j/dFny9CpYBmpMWCh+ZwwPMeT4/Yv7TtGtRTLK1SqP68Mf/0j82GHDv1GwAA//+AYLxzVQEAAA==")
	data["docstring/seq.md"] = decompress("H4sIAAAAAAAA/2yPwWqEMBiE73mKIXtoAtV36KGH3gs9iEhWRwzExE3+de3bF5UtPfQyEPLNfPwXmMIbxpRnCyfCeZECSehTXJnl+NnfDoW3O2NP9eVDeLJ/UZmIJafVDxz+6cGP8AJf4ovAhUw3fNf4GI/egfcuxiS48jnJ4RVd9KHDY3deiUy558ihVupywVvE++bmJVApwDwmxipQ0JTzLD0xhKRtqwDASKpW9pIyzOwWmDGi2VqYIhkbdKWtRbHWKvU5+QKe06f7FKNr9FRpaO4RfiNVuu1q9RMAAP//+bJX0E4BAAA=")
	data["docstring/some-first.md"] = decompress("H4sIAAAAAAAA/xzMy63CMBBG4X2q+KW7uUSYdOAK6CEy1oBHxB40Dx7dI7I6iyN9f/g36ZQy6P1QXEW7zfMBKaOXqoIXe4M1UU+VtQY7j9t05jthSXk54hIOJQ8dhnXwtqIYTGT8WsZnN0HPskVxsv1GbafpGwAA///BbWHeewAAAA==")
	data["docstring/some-last.md"] = decompress("H4sIAAAAAAAA/xyMXQrCMBAG33uKD3zRYu0NcgLvUGJYzWKTlf3x5/bSPg3MwBxwNGk0pQT6vhR30WbjeMJmWi4q+LBXWBX1qbCWYOf+GK78JMxTSvMZt3AoeWg3LJ3XBdlgIn1j7r/9CnrnNbKT7TVKvQz/AAAA//8zB9jPfgAAAA==")
	data["docstring/str.md"] = decompress("H4sIAAAAAAAA/5yPS2rzMBSF51rFSTL5Y34CfWwglG6g6ayURLaPa4Esmasrp9l9ieyGQmcd6ug87rfBv6SCLspQVVs0MUwUTUVI0AiLpOLCh3kSWmWCReB5EdFJHKA9l7frHFtM1mcmxK78jBIn17KdK3fGbDbYBzx/2mH0NAbzBeue3sc13u5wjwc8vm+Nee1dAmcjzs57CDVL+LGIY8ndUsd54IW2peBQPMmYvVdKsOom+st/xEA0NiAnlq5TUlmd0OXQqIvhyj1KbHPDG/+CRknF0FtFNdhLhZqwtec1UxNC26K+lNbDmFUpRaPsvlFXf2f9FTzuzFcAAAD//9gkDpPCAQAA")
	data["docstring/sym.md"] = decompress("H4sIAAAAAAAA/3yOQU7DQAxF9z7FV7OgWZA7AOIG7CuTcZmRHE/qcQq9PZoAEquu7K/3bb0Bx3Zb0MJHzNWu4tHAPRf7QLGoYOtT3CSh3Zb3qvTy14wsWL1eS+rw7tGE5xoZl421nIsksCVonVl/Cw3sAp5nWUPSRDQMeDK8fvGyqhABxyRnZFGtj116Nz/s+TCOncvlH37Y15HoLZcG+XmDz6IKl9jccArf5DTRdwAAAP//mE7l2gUBAAA=")
	data["docstring/take.md"] = decompress("H4sIAAAAAAAA/1SQPU8zMRCEe/+KkVK8d3pRIj4SaCko6JEoEIqcu7mcxcUO6zVJ+PXIvnxAOzs7z+5MUKn9IJqQvCLy83+NLERoT3ROooIDN/QaEbrsSPQNo3l1wwChJvGwGOz34TzMRjrtKZiV4BmCoOOOcgnrJGwKZMW18975dV7LwlbCl2vZXmComuAbq/RW2dZTPI/OGJI0/OWzQsQ+iFKgvfXFJXkcle345dVRjGnQTD1fvcsfrQiVVGDt1JjJBI8eT3u72Q40BqhadtjjX3WNG9zirq5P4gFvcyxwj4f3opVeF9jjUBvz0rsIjjEj6FhdvuVvectTNOZY1Mup+QkAAP//U27VOaQBAAA=")
	data["docstring/thread-first.md"] = decompress("H4sIAAAAAAAA/3SRPW7cMBCFe53iAVtEkhVpY8RtABcucoKU9qw0EglQMwI52p+cPiA3WSCFS4IzH9/3eED99Qf4ukXMGtfUtg3MRaYp4Uxh53zSfXEYKYQESjDHPmL2MRkoLvvKYtVbHibjhCHDBpBMD5C3B8UcI+3bFjxP9xc/Q/Z4lVsZgTky+ARRA4XMvIEw7zKaVynJcPEh4MQYVc4cjSd4MYUK48SzRv6bxsvSV9XhgFfB25XWLXBVobRwRP2Eb8cGddviuUE9DHhpmqr6ldl83YqT4iNf5Jn6Ccey8NzgpfkozpFtj4L37+89fgo0Thzz0onNOOLs007B/2ZcHNmXhEW9LFDpStTVL84wOtXEectL4mggDN2QiyJsgUZ2GjJ21lgKvZvxdP+x/j+f7mHU/XPqitWfAAAA//+mkkxa+gEAAA==")
	data["docstring/thread-last.md"] = decompress("H4sIAAAAAAAA/3SRsW7cMBBEe37FAFfkeLlIFwNuDbhwkS9Iae9JK5EAtSuQq/M5Xx+QMlwESCeBM48zwwOOP56ewPc1Y9K8lNPJw0JmGgtulDauf7rNAQOlVEAFFjhmJCoGyvO2sJh7qVoyLugrqwfJ+MWJ9gWxwCjbuqbI437hf4gdnuWjKWCBDLFA1ECpIj9AmDYZLKq0XHiPKeHKGFRunI1HRDGFCuPKk2b+DBNl7pw7HPAseLnTsiZ2DvsIFxy/4+fF43g64cHj2Pd49N653xXO97V1Ury1g121W3Dx3r+1zplty4LXS/fw+Nrhl0DzyLn6rmzGGbdYNkrxD+M9kH0rmDXKDJVzi7vEORiGoFq4uqIUzgZCf+7rVoQ10cBBU8VOmtumezse9zfr/u2E82er9tHin713fwMAAP//wyn2nv4BAAA=")
	data["docstring/to-assoc.md"] = decompress("H4sIAAAAAAAA/2SQzU7rQAyF9/MUR+niNrqlRWKXHQseAbFALMyM01jMT5jxNEmfHiWIComVZfl8ny3vsNd0R6Uki8Kf/1vYFC+ctaxt5Wi5QBMoYgsJqVwYRXO1WjObF/F+ZSwpR1IGobAi9b94iX8MRzxHLx8MHVYbRUfZ4fQtOqGv0aqkeMA0iB0gBZ6u4hfYFMaq7A4bmblUv23TQcoNw7Re9c4IpJyFvFzZQUJgJ6Tsl6Mxux0eI55mCqNnY4C94x4zXrtIgdGUsapybtDRmXF/fHhrf0IL/u27ieU8KBq/lqbdhrdXzlha8xUAAP//Hd++6FwBAAA=")
	data["docstring/to-list.md"] = decompress("H4sIAAAAAAAA/0yQTU7DQAyF93OKJ2XBjFpaAb0AC46AWFRdmBlHsZifkHFo0tOjBIhY+vn7bMsNrJb7KFVR+XPn4Ev+4kHrUo6cPVdoAWFBzJvEuBCelDMpg1BZUdp/tOSNP+A1R/lgaMeoSjnQEHD8GXBEO2avUvIe1058B6mIdJM4w5fUj8phv5oD1zGuW7STumm4Lte8MxIpD0JRbhwgKXEQUo7zwZimwXPGy0Spj2wMYAO3mHB+wCOecLq4v2w2AGAT9bBtxnm6wO4w4eTc2rmzv4pbg+1rE2ZnvgMAAP//e2v3tEYBAAA=")
	data["docstring/to-vector.md"] = decompress("H4sIAAAAAAAA/0yQwU7DMAyG73mKX9qBdIxNDF6AA4+AOCAOJnFVi8Qpjbt1e3rUToIe7f/7bMsbeCsPJw5WBlT+uW8Qip54sDqXI2vgCisg3CD3LinNTCBjJWMQKhtKu+JFV8Yeb5rkm2EdoxpppCHicBtxQDtqMCm6w7mT0EEqEl0lXRBK7kfjuFvMgeuYlj3WSf3TcJ7v+WJkMh6Eklw5QnLmKGScLnvnNhu8KF4nyn1i5wAfucXkAMBn6uFbxcf0Cb/dYsKxaZbozj/iiCc8N0tj9aj/BFPjfgMAAP//x3q+iEQBAAA=")
	data["docstring/vector.md"] = decompress("H4sIAAAAAAAA/1SOQYrrMBBE9zpFkWziwM8d/mJuMDDrHqmEBHLLqDuOffsBx7OYbb3i8a64rYzeB3If8/0+IQ6K0yBQvvCG4au2dpI/AK/SjWDjTHWDDMILwVXaU5zp0BqW0deamB74LNWQnxq9dkU1aEeqOXNQHV5ED8Gpb9U5pMF2ddnALXI5Xo7qiKL4JvzoSqgKQa7D/F9sYoYsVmrXRwjXK/4rPjaZl8YQgFtixoZLYWv9Mv0uOy5eOPhezogN+xR+AgAA//+vfXdiKQEAAA==")
	data["docstring/when.md"] = decompress("H4sIAAAAAAAA/1SQzW7qMBCF936KI1jcBHFB99K/VaUuuugTdImceEJGHcbIdkry9pUdSunKGn3njDXfEtW5J8UpkEPnw3G1qtF6dZzYqxWZQJ9WBpsowqIR336Ytw6ppytwpc2tTQSOSGFI/YRKfcK+sxJpv0YZlGVfr0s3fxVhw82WDd5Z5DqDbNuXHFiRhqBrBMov62HewWqvefYKG8EpIlAcJG2MWS7xongd7fEkZAxQOeow4k/1D/+xwx3u8YBHPNV1ocVE9YxKSDHW2NUGAKpTYE2iWIz5voYPiwuYc7X5NhKnY+MF27zor/q0zfkhkgNrTGQdfDfTbdGgpfXL3o9Uq67g4hzn7Ka5le5VJnB3OTd3s+TZvvkKAAD//6VYmc3YAQAA")
	data["docstring/with-meta.md"] = decompress("H4sIAAAAAAAA/1yQsU4zMRCEez/FKGly+fOnoKSjoKCnQxQbex2vdLc+btdJJMS7o7tEBNGOPdrvmzU2Z/Hyf2An5DoNGGgcOf3rQO4UCxvmt0RO8Irrp/CS4YVhI0fJwulaFUOkkQ49o2YcWPQIUq1OzmkHL2LITaNLVZyl7zGxt0lBUD5D1Jw0LmUv5DhR3xgz359rP0Q3xrTHa+F7LIZxqidJnCC6lBfAmm96MP5orJFtB2uxgAxkVqOQy4lhPrXobWLbh7Be40nxfKFh7DkE/JosAMAmK94u79hst7jgoeuW9PMx1Qis6K68SKXaDj3bvIYebfXVhe8AAAD//z7LZmuEAQAA")
	data["docstring/with-ns.md"] = decompress("H4sIAAAAAAAA/3SQsW7yQBCEez/FyDSg/wcpKelSUKTPCyy+MZx03rNu18G8fcQp2EmR9ttvR6PZYHuLft2rQWUg+lyGfzvwU9IkTqvAIBeJag6plo3SsTktjsDoyP23HBV+JbqszrlygY3sYh8Z1oAD3vsqLgQh06DZIalQwh2co/l/RMctpoQz0RWKM+DMPhc+i0a91Kha4ICP63JhQKFNqfZ4KEnMq7dEFvpUlOHQNJsN3hSnWYYxsWmAbWCPGW2e3GLg777t7mE89xvu+3UcYP2N+scrsB2Loh1LVMfLES3mXc38gV+PD9p8BQAA//9UbL2/qQEAAA==")
}
