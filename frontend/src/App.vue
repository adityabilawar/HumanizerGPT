<template>
  <div id="app" class="container">
    <div class="col-md-6 offset-md-3 py-5">
      <h1>Get 0% AI detected Essays or Papers</h1>

      <form v-on:submit.prevent="generateHumanizedResponse">
        <div class="form-group">
          <label for="website-input">Enter your prompt</label>
          <input type="text" class="form-control" id="website-input" v-model="userPrompt" required>
        </div>
        <div class="form-group">
          <button type="submit" class="btn btn-primary rounded-pill">Generate response</button>
        </div>
      </form>
      <div class="response-box">
        <textarea class="response-textarea" readonly v-model="humanizedResponse"></textarea>
        <button class="copy-button" @click="copyResponse">Copy</button>
      </div>
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
      userPrompt: '',
      humanizedResponse: '',
    } 
  },

  methods: {
    generateHumanizedResponse() {
      //call the Go API server instead of the screenshotapi directly using axios
      //if we called screenshotapi directly we need to use get instead of post
      //but since we are like sending request to backend(Go server)we need a post request 
      axios.post("http://localhost:3000/api/prompt", {
        text: this.userPrompt,
      }, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then((response) => {
        this.humanizedResponse = response.data;
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
        console.log(this.userPrompt);
      })
    }
  }
}
</script>

<style scoped>
#app {
  background-color: #333;
  color: #fff;
  font-family: Arial, sans-serif;
}

.container {
  background-color: #444;
  padding: 20px;
  border-radius: 40px;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 20px;
}

label {
  font-size: 16px;
}

input[type="text"] {
  background-color: #555;
  color: #fff;
  border: none;
  border-radius: 5px;
  padding: 10px;
  width: 100%;
}

.btn-primary {
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}

img {
  margin-top: 20px;
  max-width: 100%;
}
</style>