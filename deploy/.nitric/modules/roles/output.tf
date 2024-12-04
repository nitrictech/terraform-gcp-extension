# DO NOT EDIT
# This module is accessed via generated bindings based on the variables and outputs file
output "base_compute_role" {
  value = local.base_compute_role
  description = "The role ID for the Nitric base compute role"
}

output "bucket_read" {
  value = local.bucket_reader_role
  description = "The role ID for the Nitric bucket read role"
}

output "topic_publish" {
  value = local.topic_publisher_role
  description = "The role ID for the Nitric topic publish role"
}

output "bucket_write" {
  value = local.bucket_writer_role
  description = "The role ID for the Nitric bucket write role"
}

output "bucket_delete" {
  value = local.bucket_deleter_role
  description = "The role ID for the Nitric bucket delete role"
}

output "secret_access" {
  value       = local.secret_access_role
  description = "The role ID for the Nitric secrete access role"
}

output "secret_put" {
  value       = local.secret_put_role
  description = "The role ID for the Nitric secrete put role"
}

output "kv_read" {
  value       = local.kv_reader_role
  description = "The role ID for the Nitric kv read role"
}

output "kv_write" {
  value       = local.kv_writer_role
  description = "The role ID for the Nitric kv write role"
}

output "kv_delete" {
  value       = local.kv_deleter_role
  description = "The role ID for the Nitric kv write role"
}

output "queue_enqueue" {
  value       = local.queue_enqueue_role
  description = "The role ID for the Nitric queue enqueue role"
}

output "queue_dequeue" {
  value       = local.queue_dequeue_role
  description = "The role ID for the Nitric queue dequeue role"
}