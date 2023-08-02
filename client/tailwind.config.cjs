/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'division-gray': '#57595D',
        'division-blue': '#0556f3',
        'gray-1': '#2D2F34',
        'gray-2': '#1E1E1E',
        'gray-3': "#151515",
        'gray-4': "#343434",
        'gray-5': "#E6E6E6",
        'gray-6': "#4A4A4A",
        'gray-7': "#999999",
        'gray-8': "#ffffff80",
      }
    },
  },
  plugins: [],
}
