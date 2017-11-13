<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: 'app'
  }
</script>

<style lang="scss">
  body {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #dbdbdb;
    background: #000;
    margin: 0;
  }
  .flex {
    display: flex;
  }
  .column {
    flex-direction: column;
  }
  .align-center {
    align-items: center;
  }
  .align-end {
    align-items: flex-end;
  }
  .justify-center {
    justify-content: center;
  }
  .wrap {
    flex-wrap: wrap;
  }
  $flex: 9;
  @while $flex > 0 { // flex 1 - 9
    .flex-#{$flex} {
      flex: $flex;
    }
    $flex: $flex - 1;
  }

  a { text-decoration: none; }
  a:visited { text-decoration: none; color: #dbdbdb; }
  a:focus { text-decoration: none; color: #dbdbdb; }
  a:hover, a:active { text-decoration: none; color: red }

  ul {
    list-style: none;
  }

  .back {
    color: transparent;
    text-shadow: 0 0 0 #dbdbdb;
  }
  .pager {
    margin-left: 20px;
    font-size: 42pt;
  }

  /*This is bible stuff*/
  .p {
    display: inline;
  }
  sup {
    color: #daffbf;
    margin-left: 10px;
  }
  .wj {
    color: red;
  }
  .q1 {
    margin-left: 2rem;
  }
  .q2 {
    margin-left: 4rem;
  }
</style>
