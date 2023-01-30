const getBaseUrl = () => {
  return `${window.location.protocol}//${window.location.host}/api/wishes`
}

export const submitNewYearsWish = async (newYearsWish) => {
  const body = { message: newYearsWish }
  console.log(body)
  const response = await fetch(getBaseUrl(), {
    method: 'POST',
    body: JSON.stringify(body),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

export const getAllWishes = async (newYearsWish) => {
  const response = await fetch(getBaseUrl())
  return await response.json()
}
