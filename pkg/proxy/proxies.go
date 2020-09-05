package proxy

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type ProxyList []Proxy

func (ps ProxyList) Len() int {
	return len(ps)
}

func (ps ProxyList) TypeLen(t string) int {
	l := 0
	for _, p := range ps {
		if p.TypeName() == t {
			l++
		}
	}
	return l
}

var sortType = make(map[string]int)

func init() {
	sortType["ss"] = 1
	sortType["ssr"] = 2
	sortType["vmess"] = 3
	sortType["trojan"] = 4
}

func (ps ProxyList) Less(i, j int) bool {
	if ps[i].BaseInfo().Name == ps[j].BaseInfo().Name {
		return sortType[ps[i].BaseInfo().Type] < sortType[ps[j].BaseInfo().Type]
	} else {
		return ps[i].BaseInfo().Name < ps[j].BaseInfo().Name
	}
}

func (ps ProxyList) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ProxyList) Deduplication() ProxyList {
	result := make(ProxyList, 0, len(ps))
	temp := map[string]struct{}{}
	for _, item := range ps {
		if item != nil {
			if _, ok := temp[item.Identifier()]; !ok {
				temp[item.Identifier()] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

func (ps ProxyList) Sort() ProxyList {
	sort.Sort(ps)
	return ps
}

func (ps ProxyList) NameAddCounrty() ProxyList {
	num := len(ps)
	wg := &sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		ii := i
		go func() {
			defer wg.Done()
			_, country, err := geoIp.Find(ps[ii].BaseInfo().Server)
			if err != nil {
				country = "🏁 ZZ"
			}
			ps[ii].SetName(fmt.Sprintf("%s", country))
			ps[ii].SetCountry(country)
			//ps[ii].SetIP(ip)
		}()
	}
	wg.Wait()
	return ps
}

func (ps ProxyList) NameAddIndex() ProxyList {
	num := len(ps)
	for i := 0; i < num; i++ {
		ps[i].SetName(fmt.Sprintf("%s_%+02v", ps[i].BaseInfo().Name, i+1))
	}
	return ps
}

func (ps ProxyList) NameReIndex() ProxyList {
	num := len(ps)
	for i := 0; i < num; i++ {
		originName := ps[i].BaseInfo().Name
		country := strings.SplitN(originName, "_", 2)[0]
		ps[i].SetName(fmt.Sprintf("%s_%+02v", country, i+1))
	}
	return ps
}

func (ps ProxyList) NameAddTG() ProxyList {
	num := len(ps)
	for i := 0; i < num; i++ {
		ps[i].SetName(fmt.Sprintf("%s %s", ps[i].BaseInfo().Name, "@peekfun"))
	}
	return ps
}

func Deduplication(src ProxyList) ProxyList {
	result := make(ProxyList, 0, len(src))
	temp := map[string]struct{}{}
	for _, item := range src {
		if item != nil {
			if _, ok := temp[item.Identifier()]; !ok {
				temp[item.Identifier()] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

func (ps ProxyList) Clone() ProxyList {
	result := make(ProxyList, 0, len(ps))
	for _, pp := range ps {
		if pp != nil {
			result = append(result, pp.Clone())
		}
	}
	return result
}
