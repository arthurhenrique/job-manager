CREATE INDEX job_execution_object_id ON job_execution USING btree (object_id);
CREATE INDEX job_execution_status ON job_execution USING btree (status);
CREATE INDEX job_execution_sleep ON job_execution USING btree (sleep);