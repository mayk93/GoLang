# -*- coding: utf-8 -*-
from __future__ import unicode_literals

import os
import json

if os.path.exists("python_go_pipe"):
    os.remove("python_go_pipe")
    os.mkfifo("python_go_pipe")
else:
    os.mkfifo("python_go_pipe")

urls = json.dumps([{"url": "http://www.example.com"}])

pipe = os.open("python_go_pipe", os.O_WRONLY)
os.write(pipe, urls)
