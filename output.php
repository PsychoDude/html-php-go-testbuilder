<!DOCTYPE html><html lang="en"><head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<title><?php echo $title; ?></title>
	</head>
	<body>
		<div class="container">
			<h1 cms-title=""><?php echo $title; ?></h1>
		</div>
		<!-- Loop Container -->
		<div class="container" cms-loop=""><?php foreach ($cards as $card): ?>
			<!-- Loop Card -->
			<div class="card" cms-card="">
				<span cms-card-title=""><?php echo $card['title']; ?></span>
				<p cms-card-content=""><?php echo $card['content']; ?></p>
			</div>
			<!-- /loop card -->
		<?php endforeach; ?></div>
		<!-- /loop container -->
</body></html>