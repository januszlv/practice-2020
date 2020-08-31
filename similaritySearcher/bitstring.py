#!/usr/bin/python

import sys
from rdkit import Chem
from rdkit.Chem import AllChem

finName = sys.argv[1]
foutName = sys.argv[2]

fin = open(finName, 'r')
arg = fin.read()

mol = Chem.MolFromSmiles(arg)
fp = AllChem.RDKFingerprint(mol)


fout = open(foutName, 'w')

fout.write(fp.ToBitString())
fin.close()
fout.close()