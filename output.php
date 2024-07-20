<!DOCTYPE html><html lang="en"><head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<title>&lt;?php echo $title; ?&gt;</title>
	</head>
	<body>
		<div class="container">
			<h1 cms-title="">&lt;?php echo $title; ?&gt;</h1>
		</div>

		<div class="container" cms-loop="">&lt;?php foreach ($cards as $card): ?&gt;
			<div class="card" cms-card="">
				<span cms-card-title="">&lt;?php echo $card[&#39;title&#39;]; ?&gt;</span>
				<p cms-card-content="">&lt;?php echo $card[&#39;content&#39;]; ?&gt;</p>
			</div>

			

			
		&lt;?php endforeach; ?&gt;</div>
	

</body></html>