resource "vercel_deployment" "new_years_eve_deployment" {
  project_id = vercel_project.new_years_wish_frontend.id
  files = merge(data.vercel_project_directory.frontend.files, data.vercel_project_directory.backend.files)
  production = true
  project_settings = {
    output_directory = "frontend"
    root_directory = "frontend"
  }
}
