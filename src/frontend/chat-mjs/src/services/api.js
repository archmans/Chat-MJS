/**
 * API service for chat functionality
 */

const API_URL = 'http://localhost:8000/api/gpt'

/**
 * Fetch chat response from the API
 * @param {string} message - User message
 * @param {string} method - Algorithm method (kmp or bm)
 * @returns {Promise<string>} - Response text
 */
export async function fetchChatResponse(message, method) {
  try {
    const response = await fetch(API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        pertanyaan: message,
        method: method,
      }),
    })

    if (!response.ok) {
      throw new Error(`API error: ${response.status}`)
    }

    const data = await response.json()
    return data.response
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}
