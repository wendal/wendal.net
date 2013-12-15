#!/usr/bin/python
# -*- coding: UTF-8 -*-

import os
import os.path
import subprocess

for root, dirs, files in os.walk("day8") :
    for name in files :
        path = os.path.join(root, name)
        print path
        subprocess.call(["convert.exe", path, "-resize", "25%", path[:-4] + "_small.jpg"])