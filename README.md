# MemoryGameGolang
Ce code est une application web qui implémente un jeu de concentration. Le joueur se voit proposer un ensemble de boutons avec des images, dont plusieurs sont répétées. La tâche du joueur est de trouver toutes les paires d'images en double en cliquant sur les boutons.

Lorsque le bouton est cliqué, une requête AJAX est envoyée au serveur Go. La requête transmet le nom du fichier image, qui est récupéré à partir de l'URL de l'arrière-plan du bouton. Le serveur vérifie si les deux clics correspondent et renvoie une réponse AJAX, qui est traitée sur le client.

Ce code contient plusieurs fonctions :

generateButtonsHTML() : génère le code HTML pour le plateau de jeu, y compris les boutons d'image.
handleUrl() : gère les requêtes AJAX reçues lors du clic sur les boutons.
hidden() : prend en entrée le résultat d'un test en deux clics et renvoie une réponse en utilisant AJAX.
Le code implémente également l'animation de l'apparition et du masquage de l'image sur la carte lorsque vous cliquez dessus.
