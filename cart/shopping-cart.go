package cart

// Cart represents the state of a buyer's shopping cart
type Cart struct {
	items map[string]Item
}

// Item represents any item available for sale
type Item struct {
	ID    string
	Name  string
	Price float64
	Qty   int
}

//Initiates the items map- this method will be called from all other methods which need to access the items map directly
func (c *Cart) init() {
	if c.items == nil {
		c.items = map[string]Item{}
	}
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(i Item) {
	c.init()
	if existingItem, ok := c.items[i.ID]; ok {
		//here i.ID is the key of c.items map
		//if the item exists then value of key i.ID is added to existingItem and boolean valriable ok is set to be true
		existingItem.Qty++
		c.items[i.ID] = existingItem
		//storing the item back to map
	} else {
		i.Qty = 1
		c.items[i.ID] = i
		//new item is stored in c.items using key i.ID
	}
}

// RemoveItem removes n number of items with give id from the cart
func (c *Cart) RemoveItem(id string, n int) {
	c.init()
	if existingItem, ok := c.items[id]; ok {
		if existingItem.Qty <= n {
			delete(c.items, id)
		} else {
			existingItem.Qty -= n
			c.items[id] = existingItem
		}
	}
}

// TotalAmount returns the total amount of the cart
func (c *Cart) TotalAmount() float64 {
	c.init()
	totalAmount := 0.0
	for _, i := range c.items {
		totalAmount += i.Price * float64(i.Qty)
	}
	return totalAmount
}

// TotalUnits returns the total number of units across all items in the cart
func (c *Cart) TotalUnits() int {
	c.init()
	totalUnits := 0
	for _, i := range c.items {
		totalUnits += i.Qty
	}
	return totalUnits
}

// TotalUniqueItems returns the number of unique items in the cart
func (c *Cart) TotalUniqueItems() int {
	return len(c.items)
}
