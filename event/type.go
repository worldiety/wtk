// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// event contains the wtk specific event types
package event

// A Type defines which kind of event is listened to.
type Type string

const (
	// Click occurs when the user clicks on a View
	Click    Type = "click"
	KeyDown  Type = "keydown"
	KeyPress Type = "keypress"
	KeyUp    Type = "keyup"
	FocusOut Type = "focusout"
	FocusIn  Type = "focusin"
	Blur     Type = "blur"
)
