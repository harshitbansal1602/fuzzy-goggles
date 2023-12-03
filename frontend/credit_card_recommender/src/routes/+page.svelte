<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import CreditCard from './credit_card.svelte';
	let workSelectedChip = '';
	let cc = [];
	let current = 1;

	const URL = 'http://localhost:8080/top';

	const workChipData = [
		{ id: 'Student', value: 'Student' },
		{ id: 'JrEmployee', value: 'Junior Employee' },
		{ id: 'SrEmployee', value: 'Senior Employee' },
		{ id: 'CXO', value: 'CXO' }
	];

	function selectWorkChip(id) {
		workSelectedChip = id;
	}

	let lifeStyleSelectedChips = [];
	const lifeStyleChipData = [
		{ id: 'LongDrives', value: 'Love Long Drives' },
		{ id: 'Foodie', value: 'Foodie' },
		{ id: 'ImpulsiveBuyer', value: 'Implusive Buyer' },
		{ id: 'BingeWatcher', value: 'Binge Watcher' },
		{ id: 'LoveGadgets', value: 'Gadgets Lover' }
	];

	function toggleChipLifeStyle(id) {
		if (lifeStyleSelectedChips.includes(id)) {
			lifeStyleSelectedChips = lifeStyleSelectedChips.filter((chipId) => chipId !== id);
		} else {
			lifeStyleSelectedChips = [...lifeStyleSelectedChips, id];
		}
	}
	async function findCreditCards() {
		const response = await fetch(URL, {
			method: 'POST',
			body: JSON.stringify({
				work: workSelectedChip,
				lifestyle: lifeStyleSelectedChips
			})
		})
			.then((data) => {
				return data.json();
			})
			.then((data) => {
				cc = data;
				const resultsSection = document.querySelector('.results');
				// @ts-ignore
				resultsSection.style.display = 'block';
				// @ts-ignore
				resultsSection.style.transform = 'scale(1)';
			})

			.catch((err) => {
				console.log(err);
			});

		const resultsSection = document.querySelector('.results');
		// @ts-ignore
		resultsSection.style.display = 'block';
		// @ts-ignore
		resultsSection.style.transform = 'scale(1)';
	}

	let container;

	onMount(() => {
		initMultiStepForm();
	});

	function initMultiStepForm() {
		const progressNumber = container.querySelectorAll('.step').length;
		const slidePage = container.querySelector('.slide-page');
		// const submitBtn = container.querySelector('.submit');
		const progressText = container.querySelectorAll('.step p');
		const progressCheck = container.querySelectorAll('.step .check');
		const bullet = container.querySelectorAll('.step .bullet');
		const pages = container.querySelectorAll('.page');
		const nextButtons = container.querySelectorAll('.next');
		const prevButtons = container.querySelectorAll('.prev');
		const stepsNumber = pages.length;

		console.log(progressCheck);
		if (progressNumber !== stepsNumber) {
			console.warn('Error, number of steps in progress bar do not match number of pages');
		}

		for (let i = 0; i < nextButtons.length; i++) {
			nextButtons[i].addEventListener('click', function (event) {
				event.preventDefault();
				// inputsValid = validateInputs(this);
				let inputsValid = true;

				if (inputsValid) {
					slidePage.style.marginLeft = `-${(100 / stepsNumber) * current}%`;
					bullet[current - 1].classList.add('active');
					progressCheck[current - 1].classList.add('active');
					progressText[current - 1].classList.add('active');
					current += 1;
				}
			});
		}

		for (let i = 0; i < prevButtons.length; i++) {
			prevButtons[i].addEventListener('click', function (event) {
				event.preventDefault();
				slidePage.style.marginLeft = `-${(100 / stepsNumber) * (current - 2)}%`;
				bullet[current - 2].classList.remove('active');
				progressCheck[current - 2].classList.remove('active');
				progressText[current - 2].classList.remove('active');
				current -= 1;
			});
		}
		// submitBtn.addEventListener('click', function () {
		// 	bullet[current - 1].classList.add('active');
		// 	progressCheck[current - 1].classList.add('active');
		// 	progressText[current - 1].classList.add('active');
		// 	current += 1;
		// 	setTimeout(function () {
		// 		alert('Your Form Successfully Signed up');
		// 		location.reload();
		// 	}, 800);
		// });
	}
</script>

<div class="container" bind:this={container}>
	<header>Compare your Credit Card</header>
	<div class="progress-bar">
		<div class="step">
			<p>Work</p>
			<div class="bullet">
				<span>1</span>
			</div>
			<div class="check fas fa-check">✓</div>
		</div>
		<div class="step">
			<p>Lifestyle</p>
			<div class="bullet">
				<span>2</span>
			</div>
			<div class="check fas fa-check">✓</div>
		</div>

		<div class="step">
			<p>Results</p>
			<div class="bullet">
				<span>3</span>
			</div>
			<div class="check fas fa-check">✓</div>
		</div>
	</div>
	<div class="form-outer">
		<form action="#">
			<div class="page slide-page">
				<div class="title">What is job title?</div>
				<div>
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					{#each workChipData as { id, value }}
						<!-- svelte-ignore a11y-click-events-have-key-events -->
						<div
							class={`chip ${workSelectedChip == id ? 'selected' : ''}`}
							on:click={() => selectWorkChip(id)}
						>
							{value}
						</div>
					{/each}
				</div>
				<div class="field">
					<button class="firstNext next">Next</button>
				</div>
			</div>

			<div class="page">
				<div class="title">Select things you spend money on</div>
				<div>
					{#each lifeStyleChipData as { id, value }}
						<!-- svelte-ignore a11y-click-events-have-key-events -->
						<!-- svelte-ignore a11y-no-static-element-interactions -->
						<div
							class={`chip ${lifeStyleSelectedChips.includes(id) ? 'selected' : ''}`}
							on:click={() => toggleChipLifeStyle(id)}
						>
							{value}
						</div>
					{/each}

				</div>
				<div class="field btns">
					<button class="prev-1 prev">Previous</button>
					<button class="next-2 next" on:click={() => findCreditCards()}>Submit</button>
				</div>
			</div>
			<div class="page">
				<div class="title">Results</div>
				{#each cc as { creditCard, score, savings }}
					{#if Math.floor(score / 10) / 10 > 0}
						<CreditCard {...creditCard} {score} {savings} />
					{/if}
				{/each}
				<div class="field btns">
					<button class="prev-2 prev">Previous</button>
				</div>
			</div>
		</form>
	</div>
</div>

<style global>
	@import url('https://fonts.googleapis.com/css?family=Poppins:400,500,600,700&display=swap');
	* {
		margin: 0;
		padding: 0;
		outline: none;
		font-family: 'Poppins', sans-serif;
	}
	:root {
		--primary: #333;
		--secondary: #bcb8b8;
		--errorColor: red;
		--stepNumber: 3;
		--containerWidth: 600px;
		--bgColor: #333;
		--inputBorderColor: lightgray;
	}

	body {
		display: flex;
		align-items: center;
		justify-content: center;
		min-height: 100vh;
		overflow-x: hidden;
	}
	::selection {
		color: #fff;
		background: var(--primary);
	}
	.container {
		width: var(--containerWidth);
		background: #fff;
		text-align: center;
		border-radius: 5px;
		padding: 50px 35px 10px 35px;
	}
	.container header {
		font-size: 35px;
		font-weight: 600;
		margin: 0 0 30px 0;
	}
	.container .form-outer {
		width: 100%;
		overflow: hidden;
	}
	.container .form-outer form {
		display: flex;
		width: calc(100% * var(--stepNumber));
	}
	.form-outer form .page {
		width: calc(100% / var(--stepNumber));
		transition: margin-left 0.3s ease-in-out;
	}
	.form-outer form .page .title {
		text-align: left;
		font-size: 25px;
		font-weight: 500;
	}
	.form-outer form .page .field {
		width: var(--containerWidth);
		height: 45px;
		margin: 45px 0;
		display: flex;
		position: relative;
	}
	form .page .field .label {
		position: absolute;
		top: -30px;
		font-weight: 500;
	}
	form .page .field input {
		box-sizing: border-box;
		height: 100%;
		width: 100%;
		border: 1px solid var(--inputBorderColor);
		border-radius: 5px;
		padding-left: 15px;
		margin: 0 1px;
		font-size: 18px;
		transition: border-color 150ms ease;
	}
	form .page .field input .invalid-input {
		border-color: var(--errorColor);
	}
	form .page .field select {
		width: 100%;
		padding-left: 10px;
		font-size: 17px;
		font-weight: 500;
	}
	form .page .field button {
		width: 100%;
		height: calc(100% + 5px);
		border: none;
		background: var(--secondary);
		margin-top: -20px;
		border-radius: 5px;
		color: #fff;
		cursor: pointer;
		font-size: 18px;
		font-weight: 500;
		letter-spacing: 1px;
		text-transform: uppercase;
		transition: 0.5s ease;
	}
	form .page .field button:hover {
		background: #000;
	}
	form .page .btns button {
		margin-top: -20px !important;
	}
	form .page .btns button.prev {
		margin-right: 3px;
		font-size: 17px;
	}
	form .page .btns button.next {
		margin-left: 3px;
	}
	.container .progress-bar {
		display: flex;
		margin: 40px 0;
		user-select: none;
	}
	.container .progress-bar .step {
		text-align: center;
		width: 100%;
		position: relative;
	}
	.container .progress-bar .step p {
		font-weight: 500;
		font-size: 18px;
		color: #000;
		margin-bottom: 8px;
	}
	.progress-bar .step .bullet {
		height: 25px;
		width: 25px;
		border: 2px solid #000;
		display: inline-block;
		border-radius: 50%;
		position: relative;
		transition: 0.2s;
		font-weight: 500;
		font-size: 17px;
		line-height: 25px;
	}
	.progress-bar .step .bullet .active {
		border-color: var(--primary);
		background: var(--primary);
	}
	.progress-bar .step .bullet span {
		position: absolute;
		left: 50%;
		transform: translateX(-50%);
	}
	.progress-bar .step .bullet.active span {
		display: none;
	}
	.progress-bar .step .bullet:before,
	.progress-bar .step .bullet:after {
		position: absolute;
		content: '';
		bottom: 11px;
		right: -166px;
		height: 3px;
		width: 155px;
		background: #262626;
	}
	.progress-bar .step .bullet .active:after {
		background: var(--primary);
		transform: scaleX(0);
		transform-origin: left;
		animation: animate 0.3s linear forwards;
	}
	@keyframes animate {
		100% {
			transform: scaleX(1);
		}
	}
	.progress-bar .step:last-child .bullet:before,
	.progress-bar .step:last-child .bullet:after {
		display: none;
	}
	.progress-bar .step p .active {
		color: var(--primary);
		transition: 0.2s linear;
	}
	.progress-bar .step .check {
		position: absolute;
		left: 50%;
		top: 70%;
		font-size: 15px;
		transform: translate(-50%, -50%);
		display: none;
	}
	.progress-bar .step .check.active {
		display: block;
		color: #000;
	}

	.chip {
		display: inline-block;
		padding: 8px;
		margin: 4px;
		background-color: #3498db;
		color: #fff;
		border-radius: 4px;
		cursor: pointer;
	}

	.selected {
		background-color: #2ecc71;
	}

	@media screen and (max-width: 660px) {
		:root {
			--containerWidth: 400px;
		}
		.progress-bar .step p {
			display: none;
		}
		.progress-bar .step .bullet::after,
		.progress-bar .step .bullet::before {
			display: none;
		}
		.progress-bar .step .bullet {
			display: flex;
			align-items: center;
			justify-content: center;
		}
		.progress-bar .step .check {
			position: absolute;
			left: 50%;
			top: 50%;
			font-size: 15px;
			transform: translate(-50%, -50%);
			display: none;
		}
		.step {
			display: flex;
			align-items: center;
			justify-content: center;
		}
	}
	@media screen and (max-width: 490px) {
		:root {
			--containerWidth: 100%;
		}
		.container {
			box-sizing: border-box;
			border-radius: 0;
		}
	}
</style>
