# Releases

## `v1.0.0 - (Jan 13 2025)`

---

#### Bugfixes / Improvements

- Major re-write of the extension to improve performance
- Updated extension framework to use latest _@quasar/app-vite_ tool chain.

## `v0.9.2 - (Apr 12 2022)`

---

#### Bugfixes / Improvements

- Bug fix, handling bad data which prevented listing of selections in _Options_ page

## `v0.9.1 - (Apr 10 2022)`

---

- Migrating extension to chrome extension _manifest v3_.
- Updating application framework to use _@quasar/app-vite_ tool chain.

## `v0.8.1 - (Mar 21 2022)`

---

#### Bugfixes / Improvements

- On _Options_ page
  - Adding pagination in "Selected Phrase" section.
  - Fixing default icon for URL under "Selected Phrase" section.
- Fixing rendering of selected phrases on the web-page. Using JSON serialization to remove Vue3 reactivity.

## `v0.8.0 - (Mar 16 2022)`

---

#### Bugfixes / Improvements

- Updating framework to Vue 3.

## `v0.7.0 - (Apr 05 2021)`

---

#### Bugfixes / Improvements

- Adding support for multi-lingual application strings, starting the Spanish.
- On _Options_ page, adding link to release notes.

## `v0.6.1 - (Mar 27 2021)`

---

#### Bugfixes / Improvements

- Fixed a bug caused due to the new _Enable by default_ option wherein the extension is disabled on websites with selected phrases.

## `v0.6.0 - (Mar 26 2021)`

---

#### Features

- On _Options_ page, under _Settings_ we have new option to enable the extension on all web pages by default. This option can reduce any frustration caused by conflicting feature of the web page.

## `v0.5.1 - (Jan 25 2021)`

---

#### Bugfixes / Improvements

- On _Options_ page, updating sort order of the selected phrases.

## `v0.5.0 - (Dec 08 2020)`

---

#### Bugfixes / Improvements

- Re-organized _Options_ page sections.

## `v0.4.1 - (Nov 24 2020)`

---

#### Bugfixes / Improvements

- Optimizing extension initialization.

## `v0.4.0 - (Nov 19 2020)`

---

#### Features

- Added a tabbed section listing the selected phrases on the current page.
- In the new section, user can:
  - Change color of the highlight
  - Scroll to the location of the highlight directly
  - Share the phrase text as a link. This feature is based on [Text Fragments](https://wicg.github.io/scroll-to-text-fragment/) unofficial draft and is supported in _Chrome_ browser only.

## `v0.3.1 - (Oct 29 2020)`

---

#### Bugfixes / Improvements

- Fixed the issue where extension was incorrectly initialized after the webpage load.

## `v0.3.0 - (Oct 28 2020)`

---

#### Features

- Added badge to show the disabled status of the extension on the current page.
- On _Options_ page,
  - Added a section to list disabled domains with the ability to enable the listed domain or url.

#### Bugfixes / Improvements

- On _Options_ page,
  - Show the disabled status of the URLs containing one or more selections.
- Improved the CSS of the extension.
- Fixed _enable_ toggle operation for the web-page.
- Improved memory management by loading extension only when required.

## `v0.2.0 - (Oct 14 2020)`

---

#### Features

- User has the ability to disable the extension for a single page of the entire domain.

#### Bugfixes / Improvements

- Improved processing of [Open Graph](https://ogp.me/) meta tags.
- On _Options_ page,
  - Updated default placeholder website image, when no _Open Graph_ meta tags are available.
  - Fixed how the website list is re-rendered after clicking **FORGET** button.

## `v0.1.0 - (Sep 26 2020)`

---

Finally got to the initial release of Illuminate browser extension.

#### Features

- Customize color list for text highlighting.
- Clear all text highlights with single click.
- See the summary of selections on the 'Options' page.
