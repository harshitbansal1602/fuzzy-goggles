<script>
// @ts-nocheck

    
    export let name
    // @ts-ignore
    export let score
    export let savings


    // Function to calculate the color based on the value
    // @ts-ignore
    function calculateColor(value) {
      const hue = ((value) / 100) * 120; // Map the value to a hue between 0 (red) and 120 (green)
      return `background-color: hsl(${hue}, 100%, 40%)`;
    }
  </script>
<div class="card">
	<span>{name}</span>
	<span class="score score-10" style="{calculateColor(score)}">{Math.floor(score/10)}/10</span>
</div>

<style>
    .score {
		font-weight: bold;
		padding: 5px 10px;
		border-radius: 4px;
		color: black;
	}


	.card {
		margin-bottom: 10px;
		padding: 10px;
		border: 1px solid #ddd;
		border-radius: 4px;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

  </style>