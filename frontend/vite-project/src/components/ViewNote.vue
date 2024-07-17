<template>
  <div class="view-note">
    <h2>Note Details</h2>
    <div>
      <strong>Title:</strong> {{ note.title }}
    </div>
    <div>
      <strong>Content:</strong>
      <pre>{{ note.content }}</pre>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ViewNote',
  props: {
    note: {
      type: Object,
      required: true
    }
  }
};
</script>

<style scoped>
.view-note {
  max-width: 800px;
  width: 600px;
  margin: 0 auto;
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 1em;
}

.view-note div {
  margin-bottom: 1em;
}

.view-note strong {
  color: #ffffff;
}

.view-note pre {
  white-space: pre-wrap;
}
</style>
