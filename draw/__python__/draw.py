#!/usr/local/bin/python3
from rdkit import Chem, DataStructs
from rdkit.Chem import AllChem, Draw
import json
import sys

if __name__ == '__main__':
    with open('__python__/' + sys.argv[1]) as json_file:
        data = json.load(json_file)
    name = data['GroupName']
    smilesGroup = data['Group']

    molGroup = [Chem.MolFromSmiles(smiles) for smiles in smilesGroup]
    for m in molGroup: AllChem.Compute2DCoords(m)
    img = Draw.MolsToGridImage(molGroup, molsPerRow = 4, subImgSize = (200, 200))
    img.save(sys.argv[2] + '.o.png')
