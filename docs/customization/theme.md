# Theme Configuration

Relay allows you to customize its visual appearance by updating your configuration file.

## Changing the Theme

You can change the active theme color by updating the `name` field under the `[Theme]` section in your configuration file.

### Configuration Location
The configuration file is formatted in **TOML** and is stored in your system's standard configuration directory, determined by Go's [`os.UserConfigDir()`](https://pkg.go.dev/os#UserConfigDir):

* **Linux:** `~/.config/`
* **macOS:** `~/Library/Application Support/`
* **Windows:** `%AppData%`

### Available Themes
Set the `name` variable to one of the following supported options:

* `"default"` — Default theme
* `"cyberpunk"` — Cyberpunk theme
* `"light-paper"` — Light Paper theme
* `"forest"` — Forest theme

### Example (`config.toml`)

```toml
[Theme]
name = "forest"
