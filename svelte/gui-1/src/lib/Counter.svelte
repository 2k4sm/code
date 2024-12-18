<script lang="ts">
  let count: number = $state(0)
</script>

<button onclick={() => count++} class="btn btn-primary">
  count is {count}
</button>
