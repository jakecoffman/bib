import axios from 'axios'

const books = [
  {abbr: 'Gen', name: 'Genesis'},
  {abbr: 'Exod', name: 'Exodus'},
  {abbr: 'Lev', name: 'Leviticus'},
  {abbr: 'Num', name: 'Numbers'},
  {abbr: 'Deut', name: 'Deuteronomy'},
  {abbr: 'Josh', name: 'Joshua'},
  {abbr: 'Judg', name: 'Judges'},
  {abbr: 'Ruth', name: 'Ruth'},
  {abbr: '1Sam', name: '1 Samuel'},
  {abbr: '2Sam', name: '2 Samuel'},
  {abbr: '1Kgs', name: '1 Kings'},
  {abbr: '2Kgs', name: '2 Kings'},
  {abbr: '1Chr', name: '1 Chronicles'},
  {abbr: '2Chr', name: '2 Chronicles'},
  {abbr: 'Ezra', name: 'Ezra'},
  {abbr: 'Neh', name: 'Nehemiah'},
  {abbr: 'Esth', name: 'Esther'},
  {abbr: 'Job', name: 'Job'},
  {abbr: 'Ps', name: 'Psalm'},
  {abbr: 'Prov', name: 'Proverbs'},
  {abbr: 'Eccl', name: 'Ecclesiastes'},
  {abbr: 'Song', name: 'Song of Solomon'},
  {abbr: 'Isa', name: 'Isaiah'},
  {abbr: 'Jer', name: 'Jeremiah'},
  {abbr: 'Lam', name: 'Lamentations'},
  {abbr: 'Ezek', name: 'Ezekiel'},
  {abbr: 'Dan', name: 'Daniel'},
  {abbr: 'Hos', name: 'Hosea'},
  {abbr: 'Joel', name: 'Joel'},
  {abbr: 'Amos', name: 'Amos'},
  {abbr: 'Obad', name: 'Obadiah'},
  {abbr: 'Jonah', name: 'Jonah'},
  {abbr: 'Mic', name: 'Micah'},
  {abbr: 'Nah', name: 'Nahum'},
  {abbr: 'Hab', name: 'Habakkuk'},
  {abbr: 'Zeph', name: 'Zephaniah'},
  {abbr: 'Hag', name: 'Haggai'},
  {abbr: 'Zech', name: 'Zechariah'},
  {abbr: 'Mal', name: 'Malachi'},
  {abbr: 'Matt', name: 'Matthew'},
  {abbr: 'Mark', name: 'Mark'},
  {abbr: 'Luke', name: 'Luke'},
  {abbr: 'John', name: 'John'},
  {abbr: 'Acts', name: 'Acts'},
  {abbr: 'Rom', name: 'Romans'},
  {abbr: '1Cor', name: '1 Corinthians'},
  {abbr: '2Cor', name: '2 Corinthians'},
  {abbr: 'Gal', name: 'Galatians'},
  {abbr: 'Eph', name: 'Ephesians'},
  {abbr: 'Phil', name: 'Philippians'},
  {abbr: 'Col', name: 'Colossians'},
  {abbr: '1Thess', name: '1 Thessalonians'},
  {abbr: '2Thess', name: '2 Thessalonians'},
  {abbr: '1Tim', name: '1 Timothy'},
  {abbr: '2Tim', name: '2 Timothy'},
  {abbr: 'Titus', name: 'Titus'},
  {abbr: 'Phlm', name: 'Philemon'},
  {abbr: 'Heb', name: 'Hebrews'},
  {abbr: 'Jas', name: 'James'},
  {abbr: '1Pet', name: '1 Peter'},
  {abbr: '2Pet', name: '2 Peter'},
  {abbr: '1John', name: '1 John'},
  {abbr: '2John', name: '2 John'},
  {abbr: '3John', name: '3 John'},
  {abbr: 'Jude', name: 'Jude'},
  {abbr: 'Rev', name: 'Revelation'}]

const chapterCache = {}

async function fetchChapters(bookAbbr) {
  if (chapterCache[bookAbbr] !== undefined) {
    return chapterCache[bookAbbr]
  }
  const chapters = (await axios.get('/api/versions/ESV/' + bookAbbr)).data
  chapterCache[bookAbbr] = chapters
  return chapters
}

const verseCache = {}

async function fetchVerses(bookAbbr, chapter) {
  const key = bookAbbr + chapter
  if (verseCache[key] !== undefined) {
    return verseCache[key]
  }
  const verses = (await axios.get('/api/versions/ESV/' + bookAbbr + '/' + chapter)).data
  verseCache[key] = verses
  return verses
}

export default {
  books,

  fetchChapters,
  fetchVerses
}
