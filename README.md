<p align="center">
	<img src="https://github.com/brookmg/Sing/blob/master/sing_icon.png?raw=true" alt="Sing" /><br>
	<h2 align="center"> Sing : A CLI For Music </h2>
</p>

```
    This project is in it's very initial stage, so any kind of feedback and contribution is 
    highly appreciated.
```

#### `Introduction`

Sing is a simple command line interface to access and index your 
music files on your local machine. Currently it can play a random 
song when asked and hosts a simple API on port 15477 by default. 

<p align="center">
	<img src="https://github.com/brookmg/Sing/blob/master/screenshots/help.png?raw=true" alt="Sing" /><br>
</p>

#### `Features`
* As a CLI
    - Play a random song
    - Search for a song 
    - Index music files
* As a WebServer API
    - Search for a song
    - Get music detail from within the file
    
<p align="center">
	<img src="https://github.com/brookmg/Sing/blob/master/screenshots/play.png?raw=true" alt="Sing" /><br>
</p>
    
#### `Limitations`
- The indexing is done by relying on another tool ( Everything ) which 
provides an SDK which is only available for windows. 
- Only tested on Windows OS.    
- Only deals with locally stored music files
- The searching mechanism is quite basic

#### `License`

```
Copyright (C) 2019 Brook Mezgebu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```