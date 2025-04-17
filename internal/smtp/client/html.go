package client

const (
	html = `<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Ваш результат теста на профориентацию</title>
  <style>
    body {
      font-family: 'Nunito', sans-serif;
      margin: 0;
      padding: 0;
      background-color: #f9f9f9;
    }

    .email-container {
      background-color: #ffffff;
      border-radius: 12px;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
      max-width: 600px;
      margin: 2rem auto;
      padding: 2rem;
    }

    .title {
      font-size: 1.8rem;
      font-weight: 600;
      color: #333;
      text-align: center;
      margin-bottom: 1rem;
    }

    .list-title {
      font-size: 1.2rem;
      font-weight: 500;
      color: #444;
      margin-top: 1.5rem;
      margin-bottom: 1rem;
    }

    ul {
      list-style-type: none;
      padding: 0;
      margin: 0;
      font-size: 1rem;
      color: #555;
    }

    li {
      margin: 0.5rem 0;
      padding: 0.75rem;
      border-radius: 8px;
      background-color: #f0f0f0;
      transition: background-color 0.3s ease;
    }

    li:hover {
      background-color: rgba(218, 240, 20, 0.1);
    }

    .cta-button {
      display: block;
      background-color: rgba(218, 240, 20, 1);
      color: #000;
      text-align: center;
      padding: 0.75rem;
      border-radius: 8px;
      font-size: 1rem;
      font-weight: 600;
      cursor: pointer;
      text-decoration: none;
      margin-top: 1.5rem;
    }

    .cta-button:hover {
      background-color: rgba(200, 220, 20, 1);
    }

    @media (max-width: 480px) {
      .email-container {
        padding: 1.5rem;
      }

      .title {
        font-size: 1.5rem;
      }

      .list-title {
        font-size: 1rem;
      }

      ul li {
        font-size: 0.95rem;
      }
    }
  </style>
</head>
<body>
  <div class="email-container">
    <h1 class="title">Ваш результат теста на профориентацию на сайте CareerNavigator AI</h1>

    <p>Поздравляем! Ваши результаты теста: %s. Вот список профессий, которые могут вам подойти:</p>

    <div class="list-title">Рекомендуемые профессии:</div>

    <ul>
      <li>%s</li>
      <li>%s</li>
      <li>%s</li>
    </ul>

    <p>Если вам нужно больше информации, пожалуйста, не стесняйтесь связаться с нами!</p>

  </div>
</body>
</html>`
)
