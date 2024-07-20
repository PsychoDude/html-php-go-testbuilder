<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta
			name="viewport"
			content="width=device-width, initial-scale=1.0" />
		<title><?php echo 'Title' ?></title>
	</head>
	<body>
		<div class="container">
			<h1 cms-title><?php echo 'Title' ?></h1>
		</div>

		<div
			class="container"
			cms-loop>
			<div
				class="card"
				cms-card>
				<span cms-card-title><?php echo 'Card Title' ?></span>

				<p cms-card-content><?php echo 'Card Content' ?></p>
				
			</div>
		</div>
	</body>
</html>
