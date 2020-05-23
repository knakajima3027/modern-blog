<template>
<div class="index">
  <p>記事一覧</p>
  <ul>
    <li v-for="post in posts" :key="post.ID">
      タイトル : <a>{{ post.Title }}</a>
    <li>
  </ul>
</div>
</template>

<script>
export default {
  name: 'index',
  data () {
    return {
      posts: null
    }
  },
  created :function() {
    this.axios.get('/posts')
    .then(response => this.posts = response.data)
  }
}
</script>