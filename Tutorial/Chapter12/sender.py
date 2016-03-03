# -*- coding: utf-8 -*-
from __future__ import unicode_literals

import hashlib
import os

to_hash = "This value must be hashed."
hashed = hashlib.md5(to_hash).hexdigest()

if os.path.exists("python_go_pipe"):
    os.remove("python_go_pipe")
    os.mkfifo("python_go_pipe")
else:
    os.mkfifo("python_go_pipe")

pipe = os.open("python_go_pipe", os.O_WRONLY)
os.write(pipe, hashed)
