package main

type icipher interface{
	GetNewIndex(indexLetter int) int
}