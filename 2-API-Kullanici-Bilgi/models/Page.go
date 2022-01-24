//sayfalar arası veri trasnferi yaparken biz sadece veriyi değil o sayfayla ilgili verileride başka bir yere göndermek istiyoruz.

package models

type Page struct {
	ID          int
	Name        string
	Description string
	URI         string
}
