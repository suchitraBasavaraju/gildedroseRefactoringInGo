package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateItemQuality(item)
	}

}

func updateItemQuality(item *Item) {
	const AGEDBRIE = "Aged Brie"
	const BACKSTAGEPASSES = "Backstage passes to a TAFKAL80ETC concert"
	const SULFURAS = "Sulfuras, Hand of Ragnaros"

	if item.Name == AGEDBRIE || item.Name == BACKSTAGEPASSES {
		if item.Quality < 50 {
			qualityAdjustment(item, 1)
			if item.Name == BACKSTAGEPASSES {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						qualityAdjustment(item, 1)
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						qualityAdjustment(item, 1)
					}
				}
			}
		}
	} else if item.Quality > 0 {
		if item.Name != SULFURAS {
			qualityAdjustment(item, -1)
		}
	}

	if item.Name != SULFURAS {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name == AGEDBRIE {
			if item.Quality < 50 {
				qualityAdjustment(item, 1)
			}
		} else if item.Name == BACKSTAGEPASSES {
			item.Quality = item.Quality - item.Quality
		} else if item.Quality > 0 {
			if item.Name != SULFURAS {
				qualityAdjustment(item, -1)
			}
		}
	}
}

func qualityAdjustment(item *Item, adjustment int) {
	item.Quality = item.Quality + adjustment
}
