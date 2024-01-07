package weight

func KgtoLB(n KG) LB {
	return LB(n * kgToLbRatio)
}
