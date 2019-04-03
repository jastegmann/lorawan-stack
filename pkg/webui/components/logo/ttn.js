// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'

import gradient from './gradient'
import style from './logo.styl'

export default (
  <svg viewBox="0 0 783 146">
    {gradient}
    <path className={style.cloud} d="M286.7 112.8c.7 8.6-2 16.7-7.5 22.9a32.4 32.4 0 0 1-24.6 10.2H39.7a32.9 32.9 0 0 1-23-9.6 34 34 0 0 1-10.1-24.1A33 33 0 0 1 33.8 79a34.9 34.9 0 0 1 29.6-28c.9-5.6 3.6-11 7.9-15.4 6.2-6.4 14.6-10 23.2-10a30 30 0 0 1 26 15 51.8 51.8 0 0 1 49.8-37.8A51.7 51.7 0 0 1 221.4 47c13-9.7 30.3-8.8 43.3 3.4l-1.4-1.2a35.3 35.3 0 0 1 13.5 29.4 24.7 24.7 0 0 1-5.6 14.8c3.3 1.1 6 3 8.6 5.6 3.6 3.6 6.1 8.1 6.8 13.1v.6c.1-.2.1-.1.1.1zm16.1-80.7A57 57 0 0 0 271.5 11a5.4 5.4 0 0 0-6.3 4c-1.1 2.8.7 6.4 3.7 7a45.6 45.6 0 0 1 29 23.7 45.5 45.5 0 0 1 3.3 31c-.7 2.9 1.4 6 4.4 6.4 2.7.8 5.7-.8 6.6-3.4 4.1-16.1.6-34.2-9.4-47.6zm-35.7-1.6c-4.3-.8-7.6 4.5-5.7 8.2a5.8 5.8 0 0 0 4 3.1 23.9 23.9 0 0 1 17 29.4 5.4 5.4 0 0 0 4.3 6.6c2.7.8 5.9-.8 6.7-3.6 2.8-10.2.7-21.6-5.5-30.1-5-6.8-12.5-12-20.8-13.6z" />
    <path className={style.text} d="M369.9 136.4h-7.4v-32.6h7.4l13.1 20.3v-20.3h7.4v32.6H383l-13.1-20.3v20.3zm77.6-26.2h-11.3v6.3H447v6.6h-10.7v6.6h11.3v6.6h-18.7v-32.6h18.7v6.5zm35.8-6.5h22.9v6.6h-7.7v26h-7.4v-26h-7.7v-6.6h-.1zm79.2 14.5l-5.2 18.3h-8.2l-9.2-32.6h8.1l5.5 22.5h.1l5.9-22.5h5.5l5.9 22.5h.5l5.5-22.5h8.2l-9.2 32.6h-8.2l-5.2-18.3zm71.9-15c9.1 0 16.5 7.4 16.5 16.8 0 9.6-7 16.9-16.5 16.9-9.8 0-16.5-7.4-16.5-16.9 0-9.5 7.4-16.8 16.5-16.8zm0 26.5c5.9 0 8.4-4.8 8.4-9.8 0-4.7-2.7-9.8-8.4-9.8-5.8 0-8.4 4.8-8.4 9.8 0 5 2.4 9.8 8.4 9.8zm65-26c6 0 10.7 3.3 10.7 9.8 0 4.8-2.7 8.2-6.7 9.3l11.3 13.6h-9.3l-9.9-12.9v12.9H688v-32.6l11.4-.1zm-3.4 14.6c2.9 0 6-.1 6-4.4s-3.3-4.4-6-4.4h-.8v8.5h.8v.3zm73.1 18.1L757 122.6v13.9h-7.4v-32.6h7.4v13.5l11.3-13.5h9.1L764 119.3l15.1 17.2h-10v-.1zM357.6 30.8h38.1v10.9h-12.9V85h-12.4V41.6h-12.9l.1-10.8zm57.6 54.3h-12.4V30.8h12.4v21.7h17.5V30.8H445v54.3h-12.4V63.2h-17.5l.1 21.9zm71.3-43.3H468v10.7h17.9v10.9H468v10.9h18.7v10.9h-31.2V30.8h31.1v11h-.1zm23.8-11h38.1v10.9h-12.9V85h-12.4V41.6h-12.9l.1-10.8zm57.6 54.3h-12.4V30.8h12.4v21.7h17.5V30.8h12.4v54.3h-12.4V63.2h-17.5v21.9zm52.8 0h-12.4V30.8h12.4v54.3zm22.8 0h-12.4V30.8h12.4l21.4 33.5h.1V30.8h12.4v54.3H665l-21.4-33.5h-.1v33.5zm93.4-18.7a24.7 24.7 0 0 1-25 19.7c-14.6 0-26.4-12.6-26.4-28 0-15.7 11.3-28 26.4-28 9.6 0 16.1 3.6 20.9 9.2l-8.1 8.9a13.3 13.3 0 0 0-12.1-6.9c-7.6 0-13.7 7.6-13.7 16.9 0 9.2 6.2 16.6 13.7 16.6 5.9 0 11.8-3.8 11.8-10.4H712v-9.6h25v11.7h-.1zM780.6 35l-5.1 9.9s-5.4-3.7-10.9-3.7c-4.3 0-6.3 1.8-6.3 4.8s5.2 5.2 11.3 7.8c6 2.5 12.8 7.7 12.8 14.8 0 12.9-9.9 17.5-20.5 17.5-12.8 0-20.3-7.3-20.3-7.3l6.2-10.4s7.3 6 13.3 6c2.7 0 7.8-.3 7.8-5.4 0-3.8-5.8-5.8-12.2-8.9-6.6-3.2-10.3-8.2-10.3-13.9 0-10 8.9-16.6 17.6-16.6 9.6.2 16.6 5.4 16.6 5.4z" />
  </svg>
)