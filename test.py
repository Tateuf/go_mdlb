import random

fichier = open("data.txt", "a")
for i in range(500):
    fichier.write("\n   {"+str(random.randint(24, 255))+","+str(random.randint(24, 255))+","+str(random.randint(24, 255))+"}"+",")
fichier.close()