<template>
  <div id="app" class="container">
    <div class="col-md-6 offset-md-3 py-5">
      <h1>Generate a thumbnail of a website</h1>

      <form v-on:submit.prevent="makeWesbiteThumbnail">
        <div class="form-group">
          <label for="website-input">Enter the URL of the website</label>
          <input type="text" class="form-control" id="website-input" v-model="websiteUrl" required>
                </div>
        <div class="form-group">
          <button type="submit" class="btn btn-primary">Generate Thumbnail</button>
        </div>
      </form>
      <img :src="thumbnailUrl"/>
      </div>
      </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'App',
  // Your component logic goes here
  data() {
    return {
      websiteUrl: '',
      thumbnailUrl: '',
    } 
  },

  methods: {
    makeWesbiteThumbnail() {
      //call the Go API server instead of the screenshotapi directly using axios
      //if we called screenshotapi directly we need to use get instead of post
      //but since we are like sending request to backend(Go server)we need a post request 
      axios.post("http://localhost:3000/api/thumbnail", {
        url: this.websiteUrl,
      
      }) 
      .then((response) => {
        this.thumbnailUrl = response.data.screenshot;
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
        console.log(this.websiteUrl);
      })
    }
  }
}
</script>


<!-- 
<style>
/* Your component styles go here */
</style> --> -->