package main

import "github.com/tumblr/tumblr.go"

var blacklist = map[uint64]bool{
	70462685041:  true,
	40125158967:  true,
	17622637091:  true,
	60732866133:  true,
	34887404095:  true,
	23537545550:  true,
	17569890948:  true,
	17202037842:  true,
	30870710649:  true,
	31643478991:  true,
	20255489177:  true,
	17706101219:  true,
	117718807879: true,
	135473895534: true,
}

var captionReplacement = map[uint64]string{
	103839731724: "looking at a pistol target",
	26786258185:  "looking at a laughing soldier while standing in front of a ridiculously huge fish",
	3747559476:   "looking at Quebecers",
	3747212426:   "looking at Quebecers",
	4259589500:   "looking at the first flowers of spring",
	3786556950:   "looking at a historic straw hat",
	3767589908:   "NOT looking at Bev Oda",
	3745391231:   "looking at a beaker",
	58922176350:  "looking at miniature office",
	44372188969:  "looking at the camera",
	60482046202:  "looking at puppies",
	60065329726:  "looking at bananas",
	59186481660:  "looking at bags",
	3727972862:   "looking at bread",
	3726978602:   "looking at Lafayette Square in New Orleans",
	3726696717:   "looking at a pancake",
	3726672595:   "looking at a polar bear",
	3725562570:   "looking at sculptured corbel of Queen Elizabeth II",
	39303878742:  "looking at earth",
}

func filterPosts(posts []tumblr.PostInterface) []tumblr.PostInterface {
	filtered := posts[:0]
	for _, post := range posts {
		if blacklist[post.GetSelf().Id] {
			continue
		}

		filtered = append(filtered, post)
	}

	return filtered
}
