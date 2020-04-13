# CaesarGolang

## But 
Pour mettre ne pratique ce que j'ai appris online en golang, j'ai décidé de développer ce programme. Il permet de lire le contenu dans des fichiers pour ensuite les chiffrer et écrire le contenu crypté dans un nouveau fichier.
Pour vérifier que le chiffrement se passe bien, je déchiffre ce qui a été chiffrer et je l'affiche sur la console.

## Particularité
Manipulation de fichiers
Utilisation des routines en go

## Details
main.go : point d'entrée du programme
caracteres.go : toutes les opérations sur les caractères non chiffrement
icipher.go : interface pour déterminer le nouvel index de la lettre chiffré ou déchiffrer 
cipher.go : utilisé pour déterminer le nouvel index de la lettre chiffré. Implément icipher.go
decipher.go : utilisé pour déterminer le nouvel index de la lettre déchiffré. Implément icipher.go
encryption.go : tout le code pour chiffrer / déchiffrer les mots / lettres
workflow.go : gère le workflow du chiffrement / déchiffrement.

## Divers 
Etant mon premier programme en golang, j'ai encore beaucoup de choses à apprendre. Je suis donc preneur de conseils. N'hésitez pas à m'en donner sur mon twitter @TheBigPilou ou sur plouiserre@gmail.com.
